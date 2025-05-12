// src/frontend/pages/_app.js
import '../styles/global.css';  // Mengimpor global.css di sini

function MyApp({ Component, pageProps }) {
  return <Component {...pageProps} />;
}

export default MyApp;