module.exports = {
    mode: 'development',
    entry: "./src/index.tsx",
    output: {
        filename: "bundle.js",
        path: __dirname + "/dist",
    },
    devtool: "source-map",
    resolve: {
        extensions: [".ts", ".tsx", ".css", ".js"],
    },
    devServer: {
        contentBase: './dist',
    },
    module: {
        rules: [{
                test: /\.tsx?$/,
                // loader: "awesome-typescript-loader"
                loader: "ts-loader",
                options: { reportFiles: ['src/**/*.{ts,tsx}', '!src/skip.ts'] }
            },
            {
                test: /.jsx?$/,
                exclude: /node_modules/,
                use: {
                    loader: "babel-loader",
                    options: {
                        presets: ['es2015', 'react']
                    },
                },
            },
            {
                test: /\.css$/,
                use: [
                    "style-loader",
                    "css-loader",
                ]
            },
            {
                test: /\.(woff|woff2)(\?v=\d+\.\d+\.\d+)?$/,
                use: [
                    {
                        loader: 'url-loader',
                        options: {
                            limit: 10000,
                            mimetype: 'application/font-woff',
                        },
                    },
                ],
            },
            {
                test: /\.ttf(\?v=\d+\.\d+\.\d+)?$/,
                use: [
                    {
                        loader: 'url-loader',
                        options: {
                            limit: 10000,
                            mimetype: 'application/octet-stream',
                        },
                    },
                ],
            },
            {
                test: /\.eot(\?v=\d+\.\d+\.\d+)?$/,
                loader: "file-loader",
            },
            {
                test: /\.svg(\?v=\d+\.\d+\.\d+)?$/,
                use: [
                    {
                        loader: 'url-loader',
                        options: {
                            limit: 10000,
                            mimetype: 'image/svg+xml',
                        },
                    },
                ],
            },
        ],
    },
    externals: {
        react: "React",
        "react-dom": "ReactDOM",
    },
};