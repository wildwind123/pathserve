#!/usr/bin/env node

"use strict";
var _typeof = typeof Symbol === "function" && typeof Symbol.iterator === "symbol" ? function (obj) { return typeof obj; } : function (obj) { return obj && typeof Symbol === "function" && obj.constructor === Symbol && obj !== Symbol.prototype ? "symbol" : typeof obj; };

const ncp = require("ncp").ncp;
var path = require("path");
var  fs = require('fs');
var mkdirp = require('mkdirp')
var destinationFolder = path.join(__dirname, "bin");

// Mapping from Node's `process.arch` to Golang's `$GOARCH`
var ARCH_MAPPING = {
  ia32: "386",
  x64: "amd64",
  arm: "arm",
};

// Mapping between Node's `process.platform` to Golang's
var PLATFORM_MAPPING = {
  darwin: "darwin",
  linux: "linux",
  win32: "windows",
  freebsd: "freebsd",
};

function validateConfiguration(packageJson) {
  if (!packageJson.version) {
    return "'version' property must be specified";
  }

  if (!packageJson.goBinary || _typeof(packageJson.goBinary) !== "object") {
    return "'goBinary' property must be defined and be an object";
  }

  if (!packageJson.goBinary.name) {
    return "'name' property is necessary";
  }

  if (!packageJson.goBinary.path) {
    return "'path' property is necessary";
  }
}

function parsePackageJson() {
  if (!(process.arch in ARCH_MAPPING)) {
    console.error(
      "Installation is not supported for this architecture: " + process.arch
    );
    return;
  }

  if (!(process.platform in PLATFORM_MAPPING)) {
    console.error(
      "Installation is not supported for this platform: " + process.platform
    );
    return;
  }

  var packageJsonPath = path.join(".", "package.json");
  if (!fs.existsSync(packageJsonPath)) {
    console.error(
      "Unable to find package.json. " +
        "Please run this script at root of the package you want to be installed"
    );
    return;
  }

  var packageJson = JSON.parse(fs.readFileSync(packageJsonPath));
  var error = validateConfiguration(packageJson);
  if (error && error.length > 0) {
    console.error("Invalid package.json: " + error);
    return;
  }

  // We have validated the config. It exists in all its glory
  var binName = packageJson.goBinary.name;
  var binPath = packageJson.goBinary.path;
  var version = packageJson.version;
  if (version[0] === "v") version = version.substr(1); // strip the 'v' if necessary v0.0.1 => 0.0.1

  // Binary name on Windows has .exe suffix
  if (process.platform === "win32") {
    binName += ".exe";
  }

  return {
    binName: binName,
    binPath: binPath,
    version: version,
  };
}

async function install(callback) {
  var opts = parsePackageJson();
  if (!opts) return callback(INVALID_INPUT);

  console.info(
    `Copying the relevant binary for your platform ${process.platform}`
  );
  const src = `./bin/pathserve-${ARCH_MAPPING[process.arch]}-${process.platform}`;

  let fileExist = false;
  if (fs.existsSync(src)) {
    fileExist = true;
  }
  if (!fileExist) {
    throw Error(
      `os does not supported ${src}, you can try compile patheserve from source code`
    );
  }
  
  await mkdirp(destinationFolder);
  const destinationFile = destinationFolder + '/' + opts.binName;
  console.log("destination", destinationFile);
  // const cmd = `cp ${src} ${value}/${opts.binName}`.;
  ncp(src, destinationFile, function (err) {
    if (err) {
      return console.error(err);
    }
    console.log("Installed successfully!");
  });
}

install(function (err) {
  if (err) {
    console.error(err);
    process.exit(1);
  } else {
    process.exit(0);
  }
});
