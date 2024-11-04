// frontend-react/src/components/ImageImporter.tsx

// ログイン画面用の画像をインポート
// すべての.jpgファイルをインポート
const loginmodules = import.meta.glob('../assets/LoggedOut/*.jpg', { eager: true });
// { eager: true } を使うと、すべてのモジュールが即時にロードされるため、モジュールの使用時に追加の読み込み時間が発生しない利点があるが、その分時間がかかる可能性

// モジュールから画像のパスを取得
const loginimages = Object.values(loginmodules).map((mod) => (mod as { default: string }).default);

export const getRandomImage = (): string => {
    return loginimages[Math.floor(Math.random() * loginimages.length)];
};



// API実装までダミーでスライド用画像データの取得

// バイクスライド用の画像をインポート
const motoSlideModules = import.meta.glob('../assets/motoSlideImages/*.jpg', { eager: true });
const motoSlideImages = Object.values(motoSlideModules).map((mod) => (mod as { default: string }).default);

export const getMotoSlideImages = (): string[] => {
  return motoSlideImages;
};


// 人物スライド用の画像をインポート
const humanSlideModules = import.meta.glob('../assets/humanSlideImages/*.jpg', { eager: true });
const humanSlideImages = Object.values(humanSlideModules).map((mod) => (mod as { default: string }).default);

export const getHumanSlideImages = (): string[] => {
  return humanSlideImages;
};


// マネースライド用の画像をインポート
const moneySlideModules = import.meta.glob('../assets/moneySlideImages/*.jpg', { eager: true });
const moneySlideImages = Object.values(moneySlideModules).map((mod) => (mod as { default: string }).default);

export const getMoneySlideImages = (): string[] => {
  return moneySlideImages;
};

// 車スライド用の画像をインポート
const carSlideModules = import.meta.glob('../assets/carSlideImages/*.jpg', { eager: true });
const carSlideImages = Object.values(carSlideModules).map((mod) => (mod as { default: string }).default);

export const getCarSlideImages = (): string[] => {
  return carSlideImages;
};


// スノーボードスライド用の画像をインポート
const snowboradSlideModules = import.meta.glob('../assets/snowboradSlideImages/*.jpg', { eager: true });
const snowboradSlideImages = Object.values(snowboradSlideModules).map((mod) => (mod as { default: string }).default);

export const getSnowboradSlideImages = (): string[] => {
  return snowboradSlideImages;
};

// アニメスライド用の画像をインポート
const animeSlideModules = import.meta.glob('../assets/animeSlideImages/*.jpg', { eager: true });
const animeSlideImages = Object.values(animeSlideModules).map((mod) => (mod as { default: string }).default);

export const getAnimeSlideImages = (): string[] => {
  return animeSlideImages;
};