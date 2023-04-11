//@ts-check

// eslint-disable-next-line @typescript-eslint/no-var-requires
const { withNx } = require('@nrwl/next/plugins/with-nx');

const withTM = require('next-transpile-modules')([
  '@mezzanine-ui/core',
  '@mezzanine-ui/system',
  '@mezzanine-ui/react',
  '@mezzanine-ui/icons',
  '@mezzanine-ui/react-hook-form',
  '@mezzanine-ui/react-hook-form-core',
  '@nanostores/react',
]);

/**
 * @type {import('@nrwl/next/plugins/with-nx').WithNxOptions}
 **/
const nextConfig = {
  nx: {
    // Set this to true if you would like to to use SVGR
    // See: https://github.com/gregberge/svgr
    svgr: false,
  },
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'picsum.photos',
        port: '',
        pathname: '/**',
      },
      {
        protocol: 'https',
        hostname: 'unsplash.com',
        port: '',
        pathname: '/photos/**',
      },
    ],
  },

  // webpack: (webpackConfig, options) => {

  // }
};

module.exports = withNx(withTM(nextConfig));
