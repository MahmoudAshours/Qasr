# Qasr: The Smart URL Shortener
![Qasr](https://github.com/user-attachments/assets/479a9d00-bedb-4fd1-8993-1a6d9c5b5a1d)

A modern, feature-rich URL shortener built with Golang, MongoDB, and Redis. This app goes beyond traditional URL shortening by offering unique features like editable destination URLs, dynamic redirects, advanced analytics, and more. Designed for scalability and ease of use.

## 🚀 Features

- URL Shortening: Generate short, unique codes for long URLs.
- Click Tracking: Record and display the number of clicks per short link.
- Custom Short Codes: Allow users to define custom short codes (if available).
- User-Friendly Interface: Clean, responsive UI with mobile support via a Progressive Web App (PWA).
- API Support: RESTful API for programmatic access to create, manage, and track links.
- Editable Destination URLs: Update the destination URL after creation, with versioning to track changes.
- Dynamic Redirects: Redirect users based on location, device type, or time for personalized experiences (e.g., mobile users to app stores, desktop users to websites).
- Advanced Analytics: Real-time click tracking with detailed insights (e.g., traffic sources, device types) and AI-powered recommendations for optimal sharing times.
- Custom Call-to-Actions (CTAs): Overlay branded buttons or forms on destination pages to drive engagement.
- Privacy Options: Create anonymous, non-tracked links for privacy-conscious users.
- 🔐 **Secure Short Links** — Password protection, expiration dates, and access limits.
- ⚡ **Fast Redirection** — Redis-backed cache for microsecond redirects.
- 📊 **Advanced Analytics** — Click tracking, devices, OS, location, referrers, and more.
- 📅 **Time-Limited Links** — Expire after time or number of uses.
- 🧾 **UTM Parameter Auto-Injection** — Helpful for campaigns.
- 🧭 **Geolocation Redirects** — Redirect based on country or city.
- 📱 **QR Code Generation** — Branded QR codes for any short link.
- 🗂️ **User Dashboard** — Manage links, filter by tags, view stats.
- 🌍 **Multilingual UI** — Arabic, English, French (auto-detected via browser).
- 🔒 **Spam & Phishing Detection** — With LLM-powered analysis.

## Technical Stack

Backend: Golang.
Database: MongoDB for flexible, scalable storage of links and analytics data.
Caching: Redis for high-performance caching of frequently accessed short links.
Frontend: React with Tailwind CSS for a responsive, modern UI.
API: RESTful endpoints.

## 🧠 LLM/AI-Enhanced Features

> These are powered by GPT-like models (via OpenAI API or similar).

- 🧠 **Smart Slug Generation** — Readable slugs like `short.ly/go-vs-rust`
- 🧾 **Auto-Summarize Destination Content** — View description before clicking
- 🏷️ **Auto-Tagging** — Categorize links: `tech`, `health`, `finance`, etc.
- 💬 **Natural Language Interface** — Ask: *"What were my most clicked links last week?"*
- 🧑‍💻 **AI Analytics Reports** — "Your best-performing campaign was 'RamadanPromo'."
- 🧪 **Phishing Detection** — Analyze target URL and warn users
- 🧬 **Contextual Previews** — Generate teaser content for social media

## To use

```bash
git clone https://github.com/yourname/modernurl.git
cd modernurl

go run cmd/seeder/main.go

Then access http://localhost:3000

```
### 🙌 Contributions
PRs welcome! 
