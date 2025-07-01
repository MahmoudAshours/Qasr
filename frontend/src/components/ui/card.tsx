// components/ui/card.tsx
import React from "react";

interface CardProps extends React.HTMLAttributes<HTMLDivElement> { }

export const Card: React.FC<CardProps> = ({ children, className, ...props }) => {
    return (
        <div className={`bg-white shadow-md rounded-xl p-4 ${className}`} {...props}>
            {children}
        </div>
    );
};
