import { AppProps } from 'next/app';
import Head from 'next/head';
// import '@web-nx/nextjs/styles/styles.scss';
// import '@@styles/styles.scss';
import './styles.scss';

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
