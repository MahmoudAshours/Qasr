
import { useState } from 'react';
import { Button } from "../components/ui/button";
import { Input } from "../components/ui/input";
import { Card } from "../components/ui/card";
import { useToast } from "../hooks/use-toast";
import { Copy, Link, Castle } from 'lucide-react';

interface ShortenedUrl {
    id: string;
    original: string;
    shortened: string;
    createdAt: Date;
}

const Index = () => {
    const [url, setUrl] = useState('');
    const [shortenedUrls, setShortenedUrls] = useState([] as ShortenedUrl[]);
    const [isLoading, setIsLoading] = useState(false);
    const { toast } = useToast();

    const generateShortCode = () => {
        return Math.random().toString(36).substring(2, 8);
    };

    const handleShortenUrl = () => {
        if (!url) {
            toast({
                title: "يرجى إدخال رابط",
                description: "Please enter a URL to shorten",
                variant: "destructive",
            });
            return;
        }

        if (!url.startsWith('http://') && !url.startsWith('https://')) {
            toast({
                title: "رابط غير صحيح",
                description: "Please enter a valid URL starting with http:// or https://",
                variant: "destructive",
            });
            return;
        }

        setIsLoading(true);

        // Simulate API call
        setTimeout(() => {
            const shortCode = generateShortCode();
            const newShortenedUrl: ShortenedUrl = {
                id: Date.now().toString(),
                original: url,
                shortened: `qasr.ly/${shortCode}`,
                createdAt: new Date(),
            };

            setShortenedUrls(prev => [newShortenedUrl, ...prev.slice(0, 9)]);
            setUrl('');
            setIsLoading(false);

            toast({
                title: "تم اختصار الرابط بنجاح",
                description: "Your URL has been shortened successfully!",
            });
        }, 1000);
    };

    const copyToClipboard = async (text: string) => {
        try {
            await navigator.clipboard.writeText(`https://${text}`);
            toast({
                title: "تم النسخ",
                description: "Link copied to clipboard!",
            });
        } catch (err) {
            toast({
                title: "خطأ في النسخ",
                description: "Failed to copy link",
                variant: "destructive",
            });
        }
    };

    return (
        <div className="min-h-screen bg-amber-50 font-inter">
            {/* Hero Section */}
            <div className="relative overflow-hidden">
                <div className="bg-amber-900 py-20">
                    <div className="container mx-auto px-4 text-center">
                        <div className="flex justify-center mb-8">
                            <div className="p-6 bg-amber-800 rounded-full shadow-2xl transform hover:scale-110 transition-transform duration-300">
                                <Castle className="w-20 h-20 text-amber-100 animate-pulse" />
                            </div>
                        </div>
                        <h1 className="text-6xl md:text-7xl font-playfair font-bold text-amber-100 mb-6 animate-fade-in">
                            Qasr <span className="text-amber-300">قصر</span>
                        </h1>
                        <p className="text-2xl font-playfair text-amber-200 mb-4 animate-fade-in delay-150">
                            Your Palace of Short Links
                        </p>
                        <p className="text-lg text-amber-100 max-w-2xl mx-auto font-inter animate-fade-in delay-300">
                            Transform your long URLs into elegant, royal short links with the sophistication of a palace
                        </p>
                    </div>
                </div>
            </div>

            {/* Main Content */}
            <div className="container mx-auto px-4 py-16 -mt-12 relative z-10">
                {/* URL Shortener Card */}
                <Card className="max-w-2xl mx-auto p-10 shadow-2xl bg-white border-0 rounded-2xl transform hover:scale-105 transition-all duration-300 animate-scale-in">
                    <div className="flex items-center gap-4 mb-8">
                        <div className="p-2 bg-amber-800 rounded-lg animate-bounce">
                            <Link className="w-6 h-6 text-amber-100" />
                        </div>
                        <h2 className="text-3xl font-playfair font-semibold text-amber-900">Shorten Your URL</h2>
                    </div>

                    <div className="space-y-6">
                        <div className="relative">
                            <Input
                                type="url"
                                placeholder="Enter your long URL here... (e.g., https://example.com)"
                                value={url}
                                onChange={(e) => setUrl(e.target.value)}
                                className="w-full px-6 py-4 text-lg border-2 border-amber-200 focus:border-amber-500 focus:ring-amber-500/20 rounded-xl bg-amber-50 font-inter transition-all duration-300 hover:shadow-md"
                                onKeyPress={(e) => e.key === 'Enter' && handleShortenUrl()}
                            />
                        </div>

                        <Button
                            onClick={handleShortenUrl}
                            disabled={isLoading}
                            className="w-full py-4 text-lg font-semibold bg-amber-800 hover:bg-amber-900 text-amber-100 shadow-lg transition-all duration-300 transform hover:scale-105 hover:shadow-xl rounded-xl font-inter"
                        >
                            {isLoading ? (
                                <div className="flex items-center gap-3">
                                    <div className="w-5 h-5 border-2 border-amber-100/30 border-t-amber-100 rounded-full animate-spin"></div>
                                    Creating Royal Link...
                                </div>
                            ) : (
                                "✨ Create Royal Link"
                            )}
                        </Button>
                    </div>
                </Card>

                {/* Shortened URLs Display */}
                {shortenedUrls.length > 0 && (
                    <div className="max-w-4xl mx-auto mt-16 animate-fade-in">
                        <h3 className="text-3xl font-playfair font-semibold text-amber-900 mb-8 text-center">
                            Your Royal Links 👑
                        </h3>
                        <div className="space-y-4">
                            {shortenedUrls.map((item, index) => (
                                <Card key={item.id} className="p-6 bg-white border-0 hover:shadow-xl transition-all duration-300 rounded-xl transform hover:-translate-y-1 animate-slide-in-right" style={{ animationDelay: `${index * 100}ms` }}>
                                    <div className="flex items-center justify-between gap-4">
                                        <div className="flex-1 min-w-0">
                                            <div className="flex items-center gap-2 mb-3">
                                                <div className="w-2 h-2 bg-amber-500 rounded-full animate-pulse"></div>
                                                <span className="text-sm text-amber-700 font-inter">
                                                    {item.createdAt.toLocaleString()}
                                                </span>
                                            </div>
                                            <p className="text-sm text-amber-600 truncate mb-2 font-inter">
                                                Original: {item.original}
                                            </p>
                                            <div className="flex items-center gap-3">
                                                <div className="p-1 bg-amber-800 rounded transform hover:rotate-12 transition-transform duration-200">
                                                    <Link className="w-4 h-4 text-amber-100" />
                                                </div>
                                                <p className="text-lg font-semibold text-amber-800 font-inter">
                                                    https://{item.shortened}
                                                </p>
                                            </div>
                                        </div>
                                        <Button
                                            onClick={() => copyToClipboard(item.shortened)}
                                            variant="outline"
                                            className="border-amber-300 text-amber-800 hover:bg-amber-100 hover:border-amber-500 rounded-lg font-inter transition-all duration-200 hover:scale-110"
                                        >
                                            <Copy className="w-4 h-4 mr-2" />
                                            Copy
                                        </Button>
                                    </div>
                                </Card>
                            ))}
                        </div>
                    </div>
                )}

                {/* Features Section */}
                <div className="max-w-6xl mx-auto mt-24">
                    <h3 className="text-4xl font-playfair font-bold text-center text-amber-900 mb-16 animate-fade-in">
                        Why Choose Qasr? <span className="text-amber-700">قصر</span>
                    </h3>
                    <div className="grid md:grid-cols-3 gap-8">
                        <Card className="p-8 text-center bg-white border-0 hover:shadow-2xl transition-all duration-500 hover:-translate-y-4 rounded-2xl animate-fade-in hover:bg-amber-50">
                            <div className="w-16 h-16 bg-amber-800 rounded-full mx-auto mb-6 flex items-center justify-center transform hover:rotate-12 transition-transform duration-300">
                                <Castle className="w-8 h-8 text-amber-100" />
                            </div>
                            <h4 className="text-xl font-playfair font-semibold text-amber-900 mb-3">Royal Simplicity</h4>
                            <p className="text-amber-700 font-inter">Clean, elegant interface inspired by palace architecture</p>
                        </Card>

                        <Card className="p-8 text-center bg-white border-0 hover:shadow-2xl transition-all duration-500 hover:-translate-y-4 rounded-2xl animate-fade-in delay-150 hover:bg-amber-50">
                            <div className="w-16 h-16 bg-amber-700 rounded-full mx-auto mb-6 flex items-center justify-center transform hover:rotate-12 transition-transform duration-300">
                                <Link className="w-8 h-8 text-amber-100" />
                            </div>
                            <h4 className="text-xl font-playfair font-semibold text-amber-900 mb-3">Instant Creation</h4>
                            <p className="text-amber-700 font-inter">Transform long URLs into short, memorable links in seconds</p>
                        </Card>

                        <Card className="p-8 text-center bg-white border-0 hover:shadow-2xl transition-all duration-500 hover:-translate-y-4 rounded-2xl animate-fade-in delay-300 hover:bg-amber-50">
                            <div className="w-16 h-16 bg-amber-600 rounded-full mx-auto mb-6 flex items-center justify-center transform hover:rotate-12 transition-transform duration-300">
                                <Copy className="w-8 h-8 text-amber-100" />
                            </div>
                            <h4 className="text-xl font-playfair font-semibold text-amber-900 mb-3">Easy Sharing</h4>
                            <p className="text-amber-700 font-inter">One-click copying makes sharing your links effortless</p>
                        </Card>
                    </div>
                </div>
            </div>

            {/* Footer */}
            <footer className="bg-amber-900 text-amber-100 py-12 mt-24">
                <div className="container mx-auto px-4 text-center">
                    <div className="flex justify-center items-center gap-3 mb-6">
                        <div className="p-2 bg-amber-800 rounded-lg transform hover:scale-110 transition-transform duration-300">
                            <Castle className="w-6 h-6 text-amber-100" />
                        </div>
                        <span className="text-2xl font-playfair font-semibold">Qasr قصر</span>
                    </div>
                    <p className="text-amber-200 font-inter">
                        Crafting elegant short links with palace-level sophistication
                    </p>
                </div>
            </footer>
        </div>
    );
};

export default Index;
