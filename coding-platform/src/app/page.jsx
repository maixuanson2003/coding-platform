"use client";
import React, { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import {
  Code2,
  Trophy,
  Users,
  Zap,
  ChevronRight,
  Star,
  TrendingUp,
  Award,
  CheckCircle2,
} from "lucide-react";

export default function Home() {
  const [activeTab, setActiveTab] = useState(0);
  const [scrollY, setScrollY] = useState(null);
  const router = useRouter();

  useEffect(() => {
    const handleScroll = () => setScrollY(window.scrollY);

    handleScroll(); // initialize after client mounts
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  const features = [
    {
      icon: Code2,
      title: "3000+ Problems",
      desc: "From beginner to expert level",
    },
    {
      icon: Trophy,
      title: "Weekly Contests",
      desc: "Compete with developers worldwide",
    },
    {
      icon: Users,
      title: "Active Community",
      desc: "Learn from 5M+ developers",
    },
    {
      icon: Zap,
      title: "Real-time Feedback",
      desc: "Instant code execution & testing",
    },
  ];

  const stats = [
    { number: "5M+", label: "Active Users" },
    { number: "3000+", label: "Problems" },
    { number: "500K+", label: "Contest Participants" },
    { number: "30+", label: "Programming Languages" },
  ];

  // CHỖ GÂY LỖI HYDRATION → FIX = useState lazy
  const problemTypes = [
    "Algorithms",
    "Data Structures",
    "Dynamic Programming",
    "System Design",
    "Database",
    "Concurrency",
  ];

  // random number but only client-side → FIX hydration
  const randomCounts = useState(() =>
    problemTypes.map(() => Math.floor(Math.random() * 300 + 100))
  )[0];

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-950 via-slate-900 to-slate-950 text-white overflow-hidden">
      {/* Animated BG */}
      <div className="fixed inset-0 overflow-hidden pointer-events-none">
        <div className="absolute w-96 h-96 bg-blue-500/10 rounded-full blur-3xl -top-48 -left-48 animate-pulse"></div>
        <div className="absolute w-96 h-96 bg-purple-500/10 rounded-full blur-3xl top-1/2 right-0 animate-pulse delay-1000"></div>
        <div className="absolute w-96 h-96 bg-emerald-500/10 rounded-full blur-3xl bottom-0 left-1/3 animate-pulse delay-2000"></div>
      </div>

      {/* NAV */}
      <nav
        className={`fixed top-0 w-full z-50 transition-all duration-300 ${
          scrollY > 50
            ? "bg-slate-950/90 backdrop-blur-lg border-b border-slate-800"
            : ""
        }`}
      >
        <div className="max-w-7xl mx-auto px-6 py-4 flex items-center justify-between">
          <div className="flex items-center gap-2">
            <div className="w-10 h-10 bg-gradient-to-br from-blue-500 to-purple-600 rounded-lg flex items-center justify-center">
              <Code2 className="w-6 h-6" />
            </div>
            <span className="text-xl font-bold bg-gradient-to-r from-blue-400 to-purple-400 bg-clip-text text-transparent">
              CodeMaster
            </span>
          </div>

          <div className="hidden md:flex items-center gap-8">
            <a
              href="#problems"
              className="text-slate-300 hover:text-white transition"
            >
              Problems
            </a>
            <a
              href="#contests"
              className="text-slate-300 hover:text-white transition"
            >
              Contests
            </a>
            <a
              href="#community"
              className="text-slate-300 hover:text-white transition"
            >
              Community
            </a>
            <a
              href="#learn"
              className="text-slate-300 hover:text-white transition"
            >
              Learn
            </a>
          </div>

          <div className="flex items-center gap-4">
            <button
              onClick={() => {
                router.push("/login");
              }}
              className="text-slate-300 hover:text-white transition"
            >
              Sign In
            </button>
            <button className="px-6 py-2 bg-gradient-to-r from-blue-500 to-purple-600 rounded-lg hover:shadow-lg hover:shadow-blue-500/50 transition-all">
              Get Started
            </button>
          </div>
        </div>
      </nav>

      {/* HERO SECTION */}
      <section className="relative pt-32 pb-20 px-6">
        <div className="max-w-7xl mx-auto">
          <div className="text-center max-w-4xl mx-auto">
            <div className="inline-flex items-center gap-2 px-4 py-2 bg-slate-800/50 border border-slate-700 rounded-full mb-8 backdrop-blur-sm">
              <Star className="w-4 h-4 text-yellow-400" />
              <span className="text-sm text-slate-300">
                Join 5M+ developers mastering coding
              </span>
            </div>

            <h1 className="text-6xl md:text-7xl font-bold mb-6 leading-tight">
              Master{" "}
              <span className="bg-gradient-to-r from-blue-400 via-purple-400 to-pink-400 bg-clip-text text-transparent">
                Coding Skills
              </span>
              <br />
              Through Practice
            </h1>

            <p className="text-xl text-slate-400 mb-12 max-w-2xl mx-auto">
              Level up your programming skills with thousands of problems,
              weekly contests, and a vibrant community of developers
            </p>

            <div className="flex flex-col sm:flex-row gap-4 justify-center mb-16">
              <button className="group px-8 py-4 bg-gradient-to-r from-blue-500 to-purple-600 rounded-xl hover:shadow-2xl hover:shadow-blue-500/50 transition-all flex items-center justify-center gap-2 font-semibold text-lg">
                Start Practicing
                <ChevronRight className="w-5 h-5 group-hover:translate-x-1 transition-transform" />
              </button>
              <button className="px-8 py-4 bg-slate-800/50 border border-slate-700 rounded-xl hover:bg-slate-800 transition-all font-semibold text-lg backdrop-blur-sm">
                View Problems
              </button>
            </div>

            {/* STATS */}
            <div className="grid grid-cols-2 md:grid-cols-4 gap-8 max-w-4xl mx-auto">
              {stats.map((stat, i) => (
                <div key={i} className="text-center">
                  <div className="text-3xl md:text-4xl font-bold bg-gradient-to-r from-blue-400 to-purple-400 bg-clip-text text-transparent mb-2">
                    {stat.number}
                  </div>
                  <div className="text-slate-400 text-sm">{stat.label}</div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </section>

      {/* FEATURES */}
      <section className="py-20 px-6 relative">
        <div className="max-w-7xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="text-4xl md:text-5xl font-bold mb-4">
              Why Choose CodeMaster?
            </h2>
            <p className="text-slate-400 text-lg">
              Everything you need to become a better programmer
            </p>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-6">
            {features.map((feature, i) => (
              <div
                key={i}
                className="group p-6 bg-slate-800/30 border border-slate-700 rounded-2xl hover:bg-slate-800/50 hover:border-blue-500/50 transition-all backdrop-blur-sm"
              >
                <div className="w-12 h-12 bg-gradient-to-br from-blue-500 to-purple-600 rounded-xl flex items-center justify-center mb-4 group-hover:scale-110 transition-transform">
                  <feature.icon className="w-6 h-6" />
                </div>
                <h3 className="text-xl font-semibold mb-2">{feature.title}</h3>
                <p className="text-slate-400">{feature.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* PROBLEM TYPES */}
      <section className="py-20 px-6 relative">
        <div className="max-w-7xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="text-4xl md:text-5xl font-bold mb-4">
              Practice By Topic
            </h2>
            <p className="text-slate-400 text-lg">
              Master specific areas of computer science
            </p>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
            {problemTypes.map((type, i) => (
              <div
                key={i}
                className="group p-6 bg-slate-800/30 border border-slate-700 rounded-xl hover:border-blue-500 transition-all cursor-pointer backdrop-blur-sm"
              >
                <div className="flex items-center justify-between">
                  <div>
                    <h3 className="text-lg font-semibold mb-1">{type}</h3>
                    <p className="text-slate-400 text-sm">
                      {randomCounts[i]}+ problems
                    </p>
                  </div>
                  <ChevronRight className="w-5 h-5 text-slate-500 group-hover:text-blue-400 group-hover:translate-x-1 transition-all" />
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      <footer className="border-t border-slate-800 py-12 px-6">
        <div className="max-w-7xl mx-auto">
          <div className="grid md:grid-cols-4 gap-8 mb-8">
            <div>
              <div className="flex items-center gap-2 mb-4">
                <div className="w-8 h-8 bg-gradient-to-br from-blue-500 to-purple-600 rounded-lg flex items-center justify-center">
                  <Code2 className="w-5 h-5" />
                </div>
                <span className="text-lg font-bold">CodeMaster</span>
              </div>
              <p className="text-slate-400 text-sm">
                Empowering developers worldwide
              </p>
            </div>
            <div>
              <h4 className="font-semibold mb-4">Product</h4>
              <ul className="space-y-2 text-slate-400 text-sm">
                <li>
                  <a href="#" className="hover:text-white transition">
                    Problems
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white transition">
                    Contests
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white transition">
                    Discuss
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white transition">
                    Interview
                  </a>
                </li>
              </ul>
            </div>
            <div>
              <h4 className="font-semibold mb-4">Company</h4>
              <ul className="space-y-2 text-slate-400 text-sm">
                <li>
                  <a href="#" className="hover:text-white transition">
                    About
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white transition">
                    Careers
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white transition">
                    Blog
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white transition">
                    Contact
                  </a>
                </li>
              </ul>
            </div>
            <div>
              <h4 className="font-semibold mb-4">Support</h4>
              <ul className="space-y-2 text-slate-400 text-sm">
                <li>
                  <a href="#" className="hover:text-white transition">
                    Help Center
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white transition">
                    Privacy
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white transition">
                    Terms
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white transition">
                    FAQ
                  </a>
                </li>
              </ul>
            </div>
          </div>
          <div className="pt-8 border-t border-slate-800 text-center text-slate-400 text-sm">
            © 2024 CodeMaster. All rights reserved.
          </div>
        </div>
      </footer>
    </div>
  );
}
