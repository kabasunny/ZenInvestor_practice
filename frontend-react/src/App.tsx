import React, { useRef } from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import Header from './components/Header';
import Footer from './components/Footer';
import Dashboard from './pages/Dashboard';
import MarketInsights from './pages/MarketInsights';
import Portfolio from './pages/Portfolio';
import Education from './pages/Education';

function App() {
  const dashboardRef = useRef<HTMLDivElement>(null);
  const marketInsightsRef = useRef<HTMLDivElement>(null);
  const portfolioRef = useRef<HTMLDivElement>(null);
  const educationRef = useRef<HTMLDivElement>(null);

  const scrollToRef = (ref: React.RefObject<HTMLDivElement>) => {
    ref.current?.scrollIntoView({ behavior: 'smooth' });
  };

  return (
    <Router>
      <div className="flex flex-col min-h-screen bg-gray-100">
        <Header
          onDashboardClick={() => scrollToRef(dashboardRef)}
          onMarketInsightsClick={() => scrollToRef(marketInsightsRef)}
          onPortfolioClick={() => scrollToRef(portfolioRef)}
          onEducationClick={() => scrollToRef(educationRef)}
        />
        <main className="flex-grow container mx-auto px-4 py-8">
          <div ref={dashboardRef}><Dashboard /></div>
          <div ref={marketInsightsRef}><MarketInsights /></div>
          <div ref={portfolioRef}><Portfolio /></div>
          <div ref={educationRef}><Education /></div>
        </main>
        <Footer />
      </div>
    </Router>
  );
}

export default App;