//const { defineConfig } = require('@vue/cli-service')
//require('dotenv').config()


// //webpack插件
const CompressionPlugin = require("compression-webpack-plugin");


const path = require("path");

function resolve(dir) {
  return path.join(__dirname, dir)
}




module.exports ={
//  transpileDependencies: true,

  devServer:{
    host:'0.0.0.0',
    port:process.env.PORT,
    // https:false,
    open:false,
    //以上的ip和端口是我们本机的;下面为需要跨域的
    proxy:{ //配置跨域
      '/api':{
        target:process.env.VUE_APP_BASE_API,
        ws:true,
        changeOrigin:true,//允许跨域
        pathRewrite:{
          '^/api':''   //请求的时候使用这个api就可以
        }
      }
    }
  },

  configureWebpack: {
    name: "后台",
    resolve: {
      alias: {
        '@': resolve('src')
      }
    },
    plugins: [
      new CompressionPlugin({
        cache: false,                   // 不启用文件缓存
        test: /\.(js|css|html)?$/i,     // 压缩文件格式
        filename: '[path].gz[query]',   // 压缩后的文件名
        algorithm: 'gzip',              // 使用gzip压缩
        minRatio: 0.8                   // 压缩率小于1才会压缩
      }),
    ],
  },

}
