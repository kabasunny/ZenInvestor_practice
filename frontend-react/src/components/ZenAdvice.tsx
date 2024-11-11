// components/ZenAdvice.tsx
import React from 'react';
import { AlertCircle } from 'lucide-react';

const ZenAdvice: React.FC = () => (
  <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
    <h2 className="text-xl font-semibold mb-4 flex items-center">
      <AlertCircle className="mr-2" /> Zenアドバイス
    </h2>
    <p>
      長期的な忍耐が持続可能な成長につながることを忘れないでください。
    </p>
  </div>
);

export default ZenAdvice;