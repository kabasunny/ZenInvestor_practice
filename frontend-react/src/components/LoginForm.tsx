<<<<<<< HEAD
import React, { useState } from 'react';
// import { useAuth } from '../context/AuthContext';
=======
import React, { useState } from "react";
import { useAuth } from "../context/AuthContext";
>>>>>>> c22ddb56ec4640cf7d6a03e8fc452cd83b596c91

interface LoginFormProps {
  onLogin: () => void; // onLoginプロパティの型定義
}

const LoginForm: React.FC<LoginFormProps> = ({ onLogin }) => {
<<<<<<< HEAD
  // const { login } = useAuth();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // ログイン処理を実行
=======
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // ログイン処理（例：API呼び出し）を実行
    console.log("onLogin メソッドが呼ばれました");
>>>>>>> c22ddb56ec4640cf7d6a03e8fc452cd83b596c91
    onLogin();
  };

  return (
<<<<<<< HEAD
    <form onSubmit={handleSubmit} className="max-w-lg mx-auto bg-purple-100 shadow-2xl rounded px-8 pt-6 pb-8 mb-4">
=======
    <form
      onSubmit={handleSubmit}
      className="max-w-md mx-auto bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4"
    >
>>>>>>> c22ddb56ec4640cf7d6a03e8fc452cd83b596c91
      <div className="mb-4">
        <label
          htmlFor="email"
          className="block text-gray-700 text-sm font-bold mb-2"
        >
          Email:
        </label>
        <input
          type="email"
          id="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        />
      </div>
      <div className="mb-6">
        <label
          htmlFor="password"
          className="block text-gray-700 text-sm font-bold mb-2"
        >
          Password:
        </label>
        <input
          type="password"
          id="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
        />
      </div>
      <div className="flex items-center justify-between">
        <button
          type="submit"
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
        >
          Login
        </button>
      </div>
    </form>
  );
};

export default LoginForm;
