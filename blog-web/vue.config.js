const TerserPlugin = require("terser-webpack-plugin");
const CompressionPlugin = require('compression-webpack-plugin');


module.exports = {
    chainWebpack: config => {
        config.module
            .rule('images')
            .use('image-webpack-loader')
            .loader('image-webpack-loader')
            .options({
                bypassOnDebug: true
            })
            .end()
    },
    configureWebpack: {
        plugins: [
            new CompressionPlugin({
                algorithm: 'gzip',
                test: /\.js$|\.html$|\.css/,
                threshold: 1024,
                deleteOriginalAssets: false
            })
        ],
        externals: {
            vue: 'Vue',
            'vue-router': 'VueRouter',
            axios: 'axios',
        },
        optimization: {
            minimize: true,
            minimizer: [
                new TerserPlugin({
                    minify: TerserPlugin.uglifyJsMinify,
                    terserOptions: {
                        compress: true,
                    },
                }),
            ],
        },
    },
    publicPath: "/blog",
    transpileDependencies: true,
    productionSourceMap: false
}
