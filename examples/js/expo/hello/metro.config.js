// Learn more https://docs.expo.io/guides/customizing-metro
const { getDefaultConfig } = require('expo/metro-config');
const path = require('path');

// The local @gnolang/gnonative module (this example links it instead of installing from npm).
const gnonativeModule = path.resolve(__dirname, '../../../../expo');

const config = getDefaultConfig(__dirname);

// npm v7+ installs the module's peer react / react-native into ../expo/node_modules.
// Block those so only this app's copies are bundled.
config.resolver.blockList = [
  ...Array.from(config.resolver.blockList ?? []),
  new RegExp(path.resolve(gnonativeModule, 'node_modules', 'react')),
  new RegExp(path.resolve(gnonativeModule, 'node_modules', 'react-native')),
];

config.resolver.nodeModulesPaths = [
  path.resolve(__dirname, './node_modules'),
  path.resolve(gnonativeModule, 'node_modules'),
];

config.resolver.extraNodeModules = {
  '@gnolang/gnonative': gnonativeModule,
};

config.watchFolders = [gnonativeModule];

config.transformer.getTransformOptions = async () => ({
  transform: {
    experimentalImportSupport: false,
    inlineRequires: true,
  },
});

module.exports = config;
