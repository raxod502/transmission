const path = require("path");

function isProduction(argv) {
  return !argv.development;
}

module.exports = (_, argv) => ({
  devtool: isProduction(argv) ? undefined : "source-map",
  entry: {
    index: "./frontend/js/src/index.tsx",
  },
  mode: isProduction(argv) ? "production" : "development",
  module: {
    rules: [
      {
        test: /\.tsx?$/i,
        loader: "ts-loader",
      },
    ],
  },
  output: {
    path: path.resolve(__dirname, "frontend/js/out/"),
    publicPath: "/js/",
    filename: "[name].js",
  },
  resolve: {
    extensions: [".js", ".ts", ".tsx"],
  },
});
