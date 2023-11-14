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
        externals: {
            vue: 'Vue',
            vuex: "Vuex",
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
    publicPath: "./",
    transpileDependencies: true,
    productionSourceMap: false,
}
