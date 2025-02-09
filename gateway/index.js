const { ApolloGateway } = require('@apollo/gateway');
const { ApolloServer } = require('@apollo/server');
const { startStandaloneServer } = require('@apollo/server/standalone');
const fetch = require('node-fetch'); // You can use any HTTP client library

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

// Set up the Apollo Gateway
const gateway = new ApolloGateway({
  serviceList: [
    { name: 'tenants', url: 'http://localhost:8081/query' },
    { name: 'products', url: 'http://localhost:8082/query' },
  ],
  fetch: customFetcher,  // Apply the custom fetcher with the timeout
});

const server = new ApolloServer({
  gateway,
  subscriptions: false,
});

startStandaloneServer(server).then(({ url }) => {
  console.log(`Apollo Server running at ${url}`);
});

