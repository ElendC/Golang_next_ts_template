import type { NextConfig } from "next";

const nextConfig: NextConfig = {
    /* config options here */

    // Developer mode feature for highlighting potential problems in the app.
    // Identifies unsafe lifecycles, legacy API usage, and more.
    reactStrictMode: true,

    // Proxy API request
    // Don't need to hardcore the whole url everytime, just use /api
    async rewrites(){
        return [
            {
                source: '/api/:path*', // Write this when fetching
                destination: 'https://localhost:443/:path*', // Instead of this
            }
        ]
    }

};

export default nextConfig;
