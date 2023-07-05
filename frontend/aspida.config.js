const fs = require('node:fs');
const input = 'src/openapi';

if (fs.readdirSync(input)) {
  fs.rmSync(input, { recursive: true });
}

/**
 * @type {import('openapi2aspida/dist/getConfig').ConfigFile}
 **/
const config = {
  input: input,
  outputEachDir: true,
  openapi: {
    inputFile: '../pkg/interfaces/router/swagger/ui/docs/swagger.yaml',
    yaml: true,
  },
};

module.exports = {
  ...config,
};
