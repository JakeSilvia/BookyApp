'use strict'

const utils = require('./utils')
const {VueLoaderPlugin} = require('vue-loader')
const HtmlWebpackPlugin = require('html-webpack-plugin')
const MiniCssExtractPlugin = require("mini-css-extract-plugin")

let baseConfig = {
  context: utils.resolve("/"),
  output: {
    path: utils.resolve('./static'),
    chunkFilename: 'js/[name].[hash].min.js',
    filename: "js/[name].[hash].js",
    publicPath: '/'
  },
  resolve: {
    modules: [utils.resolve('node_modules')],
    extensions: ['.js', '.vue', '.json'],
    alias: {
      '@': utils.resolve('src'),
      '@node_modules': utils.resolve('node_modules'),
      '@assets': utils.resolve('assets'),
      '@mixins': utils.resolve('src/mixins'),
      '@pages': utils.resolve('src/pages'),
      '@components': utils.resolve('src/components'),
      '@static': utils.resolve('static'),
    },

  },
  module: {
    rules: [
      // {
      //   test: /\.(js|vue)$/,
      //   use: 'eslint-loader',
      //   enforce: 'pre'
      // },
      {
        test: /\.vue$/,
        use: 'vue-loader'
      },
      {
        test: /\.html$/,
        use: [{
          loader: "underscore-template-loader",
          query: {
            engine: 'lodash',
          }
        }]
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: ['@babel/preset-env']
          }
        }
      },
      {
        test: /\.css$|\.sass$|\.scss$/,
        use: [
          MiniCssExtractPlugin.loader,
          {
            loader: 'css-loader'
          },
          {
            loader: 'postcss-loader',
            options: {
              plugins: () => [
                require("precss"),
                require('autoprefixer')
              ]
            }
          },
          {
            loader: 'sass-loader'
          },
        ]
      },
      {
        test: /\.(png|jpe?g|gif|svg)(\?.*)?$/,
        use: [
          {
            loader: 'file-loader',
            options: {
              name: "[name].[ext]",
              outputPath: "img/"
            }
          }
        ]
      },
    ]
  },
  plugins: [
    new VueLoaderPlugin(),
    new HtmlWebpackPlugin({
      filename: 'index.html',
      template: 'index.html',
      inject: true
    }),
  ]
}

module.exports = baseConfig
