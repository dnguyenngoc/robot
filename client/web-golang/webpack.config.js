// Config webpack stop load file type

const HTMLWebPackPlugin = require('html-webpack-plugin');
module.exports = {  
  module:{
    rules:[
      {
        test: /\.js$|jsx/,
        exclude:/node_modules/,
        use: {
          loader:'babel-loader'
        }
      },
      {
        test: /\.(png|jpg|jpeg|gif|mp4)$/,
        loader: 'file-loader'
      },
      {
        test: /\.s?css$/,
        use: [
          'style-loader',
          'css-loader',
          'sass-loader'
        ]
      }
    ]
  },
  plugins: [
    new HTMLWebPackPlugin({
      template: './src/index.html',
      filename: './index.html'
      }
    )
  ],
  devServer: {
    historyApiFallback: true
  }
};