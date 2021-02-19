'use strict'

const fs = require('fs')
const webpack = require('webpack')
const merge = require('webpack-merge')
const baseConfig = require('./webpack.config.base')
const UglifyJsPlugin = require("uglifyjs-webpack-plugin")
const WebpackCleanupPlugin = require('webpack-cleanup-plugin')
const MiniCssExtractPlugin  = require('mini-css-extract-plugin')
const OptimizeCSSAssetsPlugin = require("optimize-css-assets-webpack-plugin")

// the path(s) that should be cleaned
let pathsToClean = [
  './static'
];
let cleanOptions = {
  root:     __dirname,
  verbose:  true,
  dry:      false
};

pathsToClean.forEach(function (path) {
  if(!fs.existsSync(path)){
    fs.mkdirSync(path)
  }
});

module.exports = merge(baseConfig, {
  mode: 'production',
  devtool: false,
  optimization: {
    removeAvailableModules: true,
    removeEmptyChunks: true,
    mergeDuplicateChunks: true,
    minimizer: [
      new UglifyJsPlugin({
        cache: true,
        parallel: true,
        extractComments: true,
        sourceMap: false
      }),
      new OptimizeCSSAssetsPlugin({})
    ],
    splitChunks: {
      cacheGroups: {
        vendors: {
          name: 'chunk-vendors',
          test: /[\\/]node_modules[\\/]/,
          priority: -10,
          chunks: 'initial'
        },
        common: {
          name: 'chunk-common',
          minChunks: 2,
          priority: -20,
          chunks: 'initial',
          reuseExistingChunk: true
        }
      }
    }
  },
  plugins: [
    new WebpackCleanupPlugin(pathsToClean, cleanOptions),
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': '"production"'
    }),
    new MiniCssExtractPlugin({
      filename: 'css/[name].[chunkhash].css',
      sourceMap: false
    })
  ]
})
