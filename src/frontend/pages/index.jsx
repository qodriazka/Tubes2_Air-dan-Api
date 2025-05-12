import React, { useState } from 'react';
import Head from 'next/head';
import RecipeForm from "../components/RecipeForm";

export default function Home() {
  const [result, setResult] = useState(null);

  return (
    <>
      <Head>
        <title>Little Alchemy 2 Recipe Finder</title>
      </Head>

      <div className='container'>
        <header>
          <h1>Little Alchemy 2 Recipe Finder</h1>
          <p><strong>💧 Air dan Api 🔥</strong></p>
        </header>

        {/* Pass the setResult function to RecipeForm, which will handle the result */}
        <RecipeForm setResult={setResult} />

        {/* Only render RecipeTree and StatsPanel if result is available */}
        {result && (
          <div className="result-section">
            {/* RecipeTree and StatsPanel are handled by RecipeForm itself */}
          </div>
        )}
      </div>

      <style jsx>{`
        .container {
          max-width: 1200px;
          margin: 0 auto;
          padding: 40px;
          font-family: Arial, sans-serif;
        }

        header {
          text-align: center;
          margin-bottom: 30px;
        }

        header h1 {
          font-size: 2.5rem;
          color: #C326A4;
        }
      `}</style>
    </>
  );
}