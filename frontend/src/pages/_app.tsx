import { ApolloProvider } from "@apollo/client";
import { AppProps } from "next/app";
import apolloClient from "@/lib/apolloClient";
import "@/styles/globals.css";

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ApolloProvider client={apolloClient}>
      <Component {...pageProps} />
    </ApolloProvider>
  );
}

export default MyApp;
