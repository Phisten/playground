import { AppProps } from 'next/app';
import Head from 'next/head';
import './styles.scss';
import LogRocket from 'logrocket';

LogRocket.init('vcbyma/logrocket');
// LogRocket.identify('THE_USER_ID_IN_YOUR_APP', {
//   name: 'James Morrison',
//   email: 'jamesmorrison@example.com',
// });

function CustomApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <Head>
        <title>Welcome to nextjs!</title>
      </Head>
      <main className="app">
        <Component {...pageProps} />
      </main>
    </>
  );
}

export default CustomApp;
