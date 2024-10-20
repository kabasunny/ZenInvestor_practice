// ESLint設定のインポート
import js from '@eslint/js';
import globals from 'globals';
import reactHooks from 'eslint-plugin-react-hooks';
import reactRefresh from 'eslint-plugin-react-refresh';
import tseslint from 'typescript-eslint';

// ESLint設定をエクスポート
export default tseslint.config(
  // 特定のパスを無視する設定（ここでは 'dist' フォルダ）
  { ignores: ['dist'] },
  {
    // 使用するESLintの基本設定を拡張
    extends: [js.configs.recommended, ...tseslint.configs.recommended],
    
    // 対象ファイルの指定（TypeScriptファイル）
    files: ['**/*.{ts,tsx}'],
    
    // 言語オプションの設定
    languageOptions: {
      ecmaVersion: 2020, // 使用するECMAScriptのバージョン
      globals: globals.browser, // グローバル変数の設定
    },
    
    // プラグインの設定
    plugins: {
      'react-hooks': reactHooks, // React Hooksに関するプラグイン
      'react-refresh': reactRefresh, // React Refreshに関するプラグイン
    },
    
    // ルールの設定
    rules: {
      ...reactHooks.configs.recommended.rules, // React Hooksに関する推奨ルール
      'react-refresh/only-export-components': [
        'warn',
        { allowConstantExport: true }, // 定数エクスポートを許可するオプション
      ],
    },
  }
);
