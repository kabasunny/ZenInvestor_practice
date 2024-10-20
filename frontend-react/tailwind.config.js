/** @type {import('tailwindcss').Config} */
// Tailwind CSSの設定ファイルをエクスポート
export default {
  // Tailwind CSSが適用されるファイルを指定
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],

  // カスタムテーマ設定
  theme: {
    // テーマの拡張を行うためのフィールド
    extend: {},
  },

  // 使用するプラグインのリスト
  plugins: [],
};
