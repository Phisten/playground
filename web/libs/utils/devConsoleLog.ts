declare const DEV_CONSOLE_LOG: boolean;

// config.plugins.push(
//   new options.webpack.DefinePlugin({
//     DEV_CONSOLE_LOG: process.env.NODE_ENV === 'development',
//   }),
// );

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const devConsoleLog = (message?: any, ...optionalParams: any[]) =>
  DEV_CONSOLE_LOG && console.log(message, ...optionalParams);
// export const devconsole = {
//   log: devConsoleLog,
// };
