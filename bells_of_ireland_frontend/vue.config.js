module.exports = {
    assetsDir: "static",
    devServer: {
        port: 8081,
        proxy: {
            '/api/v1': {
              target: 'http://127.0.0.1:9000',
              changeOrigin: true,
            }
        }
    }
  }