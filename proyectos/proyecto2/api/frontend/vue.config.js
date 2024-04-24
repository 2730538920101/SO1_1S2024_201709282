const { defineConfig } = require('@vue/cli-service');

module.exports = defineConfig({
  transpileDependencies: true,
  chainWebpack: config => {
    config.plugin('define').tap(definitions => {
      definitions[0]['process.env'] = {
        ...definitions[0]['process.env'],
        VUE_APP_API_URL: JSON.stringify(process.env.VUE_APP_API_URL)
      };
      return definitions;
    });
  },
});
