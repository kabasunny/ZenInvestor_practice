// frontend-react/src/components/Modal.tsx

import React from "react";

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  children: React.ReactNode;
}

const Modal: React.FC<ModalProps> = ({ isOpen, onClose, children }) => {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-gray-600 bg-opacity-75 flex justify-center items-center z-50">
      <div className="relative bg-white p-4 rounded-lg shadow-lg w-11/12 md:w-1/2 h-auto max-h-[80vh] overflow-y-auto">
        <button
          className="sticky top-0 bg-red-500 text-white p-2 rounded-full z-50" // `sticky`スタイルを追加
          onClick={onClose}
          style={{ float: "right", marginRight: "4px", marginTop: "4px" }} // 位置を右上に調整
        >
          X
        </button>
        {children}
      </div>
    </div>
  );
};

export default Modal;
