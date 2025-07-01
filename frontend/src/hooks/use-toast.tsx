// hooks/use-toast.ts
type ToastVariant = "default" | "destructive" | "success";

interface ToastProps {
    title: string;
    description?: string;
    variant?: ToastVariant;
}

export function useToast() {
    const toast = (props: ToastProps) => {
        const { title, description, variant = "default" } = props;

        // For now, fallback to simple alert (replace with real toast system later)
        alert(`[${variant.toUpperCase()}] ${title}\n${description || ""}`);
    };

    return { toast };
}
