const CracoEsbuildPlugin = require("craco-esbuild");

module.exports = {
  plugins: [
    {
      plugin: CracoEsbuildPlugin,
      options: {
        esbuildLoaderOptions: {
          loader: "tsx",
        }
      }
    },
  ],
};
