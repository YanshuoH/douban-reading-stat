import path from 'path'; // Join paths with the right type of slash
import webpack from 'webpack'
import ExtractTextPlugin from 'extract-text-webpack-plugin'
let config = {
  entry: path.join(__dirname, 'webpack', 'index.js'), // We'll create this file later, when we write the frontend code
  output: {
    path: path.join(__dirname, 'public'),
    publicPath: '/public/',
    filename: 'bundle.js'
  },
  module: {
    loaders: [
      {
        test: /\.js$/, // Transpile all .js files from ES6 to ES5
        loaders: ['babel-loader']
      },
      {
        test: /\.css$/, // Use the style-loader for all .css files
        loader: ExtractTextPlugin.extract("style-loader", "css-loader")
      },
      {
        test: /\.(ttf|eot|svg|woff(?:2)?)$/, // Use the file-loader for fonts
        loaders: ['file-loader']
      },
      {
        test: /\.(jpg|png)$/,
        loaders: ['file-loader?img/[name].[ext]']
      },
      {
        test: /\.(ttf|eot|svg|woff(?:2)?)(\?[a-z0-9=.]+)$/,
        loaders: ['url-loader']
      }
    ]
  },
  plugins: [
    new webpack.ProvidePlugin({
      $: "jquery",
      jQuery: "jquery"
    }),
    new ExtractTextPlugin("bundle.css"),
    new webpack.optimize.UglifyJsPlugin({
      compress: { warnings: false }
    })
  ]
};

export default config;
