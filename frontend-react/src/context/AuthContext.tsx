// うまく機能しないので使わない

import React, { createContext, useContext, useState, ReactNode } from "react";

interface AuthContextType {
  isLoggedIn: boolean;
  login: () => void;
  logout: () => void;
}

// デフォルト値として空関数を設定したオブジェクトを提供
const AuthContext = createContext<AuthContextType>({
  isLoggedIn: false, // アプリケーションの初期ログイン状態に合わせてtrueまたはfalseを設定
  login: () => {},
  logout: () => {},
});
const AuthProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const login = () => {
    console.log("login メソッドが呼ばれました");
    setIsLoggedIn(true); // 状態を true に更新
    console.log("isLoggedIn:", true); // 状態変更後の確認
  };
  const logout = () => {
    console.log("login メソッドが呼ばれました");
    setIsLoggedIn(false);
  };

  return (
    <AuthContext.Provider value={{ isLoggedIn, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};

export { AuthProvider, useAuth };

// AuthContext は AuthContextType を保持する
// useContext(AuthContext) を呼び出すと、AuthContextType が返される
// AuthContextType にはログイン状態（isLoggedIn）と、ログイン/ログアウトを行うための関数（login, logout）が含まれている
// AuthProvider でラップされていないコンポーネントで useContext(AuthContext) を呼び出すと、エラーが発生する
