const { ApolloGateway, IntrospectAndCompose } = require('@apollo/gateway');
const { ApolloServer } = require('@apollo/server');
const { startStandaloneServer } = require('@apollo/server/standalone');
const fetch = require('node-fetch'); // You can use any HTTP client library

const tenantURL = process.env.GQL_TENANT_URL || 'http://localhost:8081/query'
console.log(`connecting to tenant ${tenantURL}`)

const productURL = process.env.GQL_PRODUCT_URL || 'http://localhost:8082/query'
console.log(`connecting to product ${productURL}`)

// Define your custom fetcher with a timeout
const customFetcher = async (url, options) => {
  const timeout = 5000; // Timeout in milliseconds (e.g., 5 seconds)
  
  const controller = new AbortController();
  const timeoutId = setTimeout(() => controller.abort(), timeout);

  try {
    const response = await fetch(url, { ...options, signal: controller.signal });
    return response;
  } catch (error) {
    if (error.name === 'AbortError') {
      throw new Error(`Request to ${url} timed out`);
    }
    throw error;
  } finally {
    clearTimeout(timeoutId);
  }
};

const gateway = new ApolloGateway({
  supergraphSdl: new IntrospectAndCompose({
    subgraphs: [
      { name: 'tenants', url: tenantURL },
      { name: 'products', url: productURL },
    ],
    fetch: customFetcher,
  }),
});

const server = new ApolloServer({
  gateway,
  subscriptions: false,
});

startStandaloneServer(server).then(({ url }) => {
  console.log(`Apollo Server running at ${url}`);
});

