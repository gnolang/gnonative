// Learn more https://docs.expo.io/guides/customizing-metro
const { getDefaultConfig } = require('expo/metro-config');
const { mergeConfig } = require('metro-config');
const path = require('path');

/** @type {import('expo/metro-config').MetroConfig} */
// const config = getDefaultConfig(__dirname);
//
// module.exports = config;
const config = {
  watchFolders: [path.resolve(__dirname + '/../../../api')],
};

module.exports = mergeConfig(getDefaultConfig(__dirname), config);
