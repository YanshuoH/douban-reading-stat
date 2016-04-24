import path from 'path'; // Join paths with the right type of slash

let config = {
  entry: path.join(__dirname, 'src/webpack', 'index.js'), // We'll create this file later, when we write the frontend code
  output: {
    path: path.join(__dirname, 'src/public'),
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
        loaders: ['style', 'css']
      },
      {
        test: /\.(ttf|eot|svg|woff(2)?)(\?[a-z0-9]+)?$/, // Use the file-loader for fonts
        loaders: ['file-loader']
      }
    ]
  }
};

export default config;
