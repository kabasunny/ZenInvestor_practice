import React from 'react';
import { BookOpen, Video, Award } from 'lucide-react';

const Education: React.FC = () => {
  return (
    <div className="space-y-6 mb-32">
      <h1 className="text-3xl font-bold">投資学習</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <BookOpen className="mr-2" /> 学習リソース
          </h2>
          <p>投資戦略と市場理解に関する厳選された記事やガイドにアクセスできます。</p>
          {/* 教育コンテンツのAPIコール */}
        </div>
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <Video className="mr-2" /> ウェビナーとチュートリアル
          </h2>
          <p>様々な投資トピックに関する専門家主導のセッションを視聴できます。</p>
          {/* ウェビナーとチュートリアルデータのAPIコール */}
        </div>
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg col-span-full">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <Award className="mr-2" /> 投資における感情知性
          </h2>
          <p>感情を管理し、合理的な投資判断を行うためのテクニックを学びます。</p>
          {/* 感情知性コンテンツのAPIコール */}
        </div>
      </div>
    </div>
  );
};

export default Education;