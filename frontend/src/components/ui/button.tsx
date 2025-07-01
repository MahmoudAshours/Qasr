import * as React from "react";
import type { ButtonHTMLAttributes } from "react"; // ✅ FIX

export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "default" | "outline" | "ghost";
}

export const Button: React.FC<ButtonProps> = ({ variant = "default", className, ...props }) => {
  const base = "px-4 py-2 rounded-lg font-medium transition-all focus:outline-none";
  const variants = {
    default: "bg-[#A67C52] text-white hover:bg-[#906541]",
    outline: "border border-gray-300 text-[#2D1B0A] hover:bg-gray-100",
    ghost: "bg-transparent hover:bg-gray-100 text-[#2D1B0A]",
  };

  return (
    <button className={`${base} ${variants[variant]} ${className}`} {...props} />
  );
};
