// frontend-react/tailwind.config.js

/** @type {import('tailwindcss').Config} */
// Tailwind CSSの設定ファイルをエクスポート
export default {
  // Tailwind CSSが適用されるファイルを指定
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],

  // カスタムテーマ設定
  theme: {
    // テーマの拡張を行うためのフィールド
    extend: {
      keyframes: {
        'slide-left': {
          '0%': { transform: 'translateX(100%)' },
          '100%': { transform: 'translateX(-100% * calc(var(--slide-count) - 1))' }, // 計算で移動距離を調整
        },
        'slide-right': {
          '0%': { transform: 'translateX(-100% * calc(var(--slide-count) - 1))' }, // 計算で移動距離を調整
          '100%': { transform: 'translateX(100%)' },
        },
      },
      animation: { // 左右のスライドを実現
        'slide-left': 'slide-left 20s linear infinite',
        'slide-right': 'slide-right 10s linear infinite',
      },
    },
  },

  // 使用するプラグインのリスト
  plugins: [],
};
