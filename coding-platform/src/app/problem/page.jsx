"use client";
import React, { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import {
  Search,
  Filter,
  CheckCircle2,
  Clock,
  Zap,
  Code2,
  ChevronRight,
  Star,
} from "lucide-react";

export default function Problem() {
  const [problems, setProblems] = useState([]);
  const [filteredProblems, setFilteredProblems] = useState([]);
  const [searchTerm, setSearchTerm] = useState("");
  const [selectedDifficulty, setSelectedDifficulty] = useState("all");
  const [selectedCategory, setSelectedCategory] = useState("all");
  const [showFilters, setShowFilters] = useState(false);
  const router = useRouter();

  // Mock data - thay thế bằng API call thực tế
  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(
          `${process.env.NEXT_PUBLIC_BASE_URL_API}/api/problem`,
          {
            method: "GET",
          }
        );
        const res = await response.json();
        console.log(res?.data);

        setProblems(res?.data);
        setFilteredProblems(res?.data);
      } catch (err) {
        console.log(err);
      }
    };
    fetchData();
  }, []);

  // Filter problems
  useEffect(() => {
    let filtered = problems;

    if (searchTerm) {
      filtered = filtered.filter(
        (item) =>
          item.problem.Title.toLowerCase().includes(searchTerm.toLowerCase()) ||
          item.problem.Category.toLowerCase().includes(searchTerm.toLowerCase())
      );
    }

    if (selectedDifficulty !== "all") {
      filtered = filtered.filter(
        (item) =>
          item.problem.Difficult.toLowerCase() ===
          selectedDifficulty.toLowerCase()
      );
    }

    if (selectedCategory !== "all") {
      filtered = filtered.filter(
        (item) =>
          item.problem.Category.toLowerCase() === selectedCategory.toLowerCase()
      );
    }

    setFilteredProblems(filtered);
  }, [searchTerm, selectedDifficulty, selectedCategory, problems]);

  const getDifficultyColor = (difficulty) => {
    switch (difficulty.toLowerCase()) {
      case "easy":
        return "text-green-400 bg-green-400/10 border-green-400/20";
      case "medium":
        return "text-yellow-400 bg-yellow-400/10 border-yellow-400/20";
      case "hard":
        return "text-red-400 bg-red-400/10 border-red-400/20";
      default:
        return "text-slate-400 bg-slate-400/10 border-slate-400/20";
    }
  };

  const categories = [
    "all",
    ...new Set(problems.map((p) => p.problem.Category)),
  ];
  const difficulties = ["all", "Easy", "Medium", "Hard"];

  const stats = {
    total: problems.length,
    solved: problems.filter((p) => p.isAccept).length,
    easy: problems.filter((p) => p.problem.Difficult === "Easy").length,
    medium: problems.filter((p) => p.problem.Difficult === "Medium").length,
    hard: problems.filter((p) => p.problem.Difficult === "Hard").length,
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-950 via-slate-900 to-slate-950 text-white">
      {/* Header */}
      <div className="border-b border-slate-800 bg-slate-900/50 backdrop-blur-lg sticky top-0 z-40">
        <div className="max-w-7xl mx-auto px-6 py-4">
          <div className="flex items-center justify-between mb-4">
            <div className="flex items-center gap-3">
              <div className="w-10 h-10 bg-gradient-to-br from-blue-500 to-purple-600 rounded-lg flex items-center justify-center">
                <Code2 className="w-6 h-6" />
              </div>
              <h1 className="text-2xl font-bold">Problem Set</h1>
            </div>
          </div>

          {/* Stats */}
          <div className="grid grid-cols-2 md:grid-cols-5 gap-4">
            <div className="bg-slate-800/50 rounded-lg p-3 border border-slate-700">
              <div className="text-slate-400 text-xs mb-1">Total</div>
              <div className="text-2xl font-bold text-blue-400">
                {stats.total}
              </div>
            </div>
            <div className="bg-slate-800/50 rounded-lg p-3 border border-slate-700">
              <div className="text-slate-400 text-xs mb-1">Solved</div>
              <div className="text-2xl font-bold text-green-400">
                {stats.solved}
              </div>
            </div>
            <div className="bg-slate-800/50 rounded-lg p-3 border border-slate-700">
              <div className="text-slate-400 text-xs mb-1">Easy</div>
              <div className="text-2xl font-bold text-green-400">
                {stats.easy}
              </div>
            </div>
            <div className="bg-slate-800/50 rounded-lg p-3 border border-slate-700">
              <div className="text-slate-400 text-xs mb-1">Medium</div>
              <div className="text-2xl font-bold text-yellow-400">
                {stats.medium}
              </div>
            </div>
            <div className="bg-slate-800/50 rounded-lg p-3 border border-slate-700">
              <div className="text-slate-400 text-xs mb-1">Hard</div>
              <div className="text-2xl font-bold text-red-400">
                {stats.hard}
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="max-w-7xl mx-auto px-6 py-8">
        {/* Search and Filters */}
        <div className="mb-6 space-y-4">
          <div className="flex flex-col md:flex-row gap-4">
            {/* Search */}
            <div className="flex-1 relative">
              <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-500" />
              <input
                type="text"
                placeholder="Search problems by title or category..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                className="w-full pl-11 pr-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl text-white placeholder-slate-500 focus:outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-all"
              />
            </div>

            {/* Filter Button */}
            <button
              onClick={() => setShowFilters(!showFilters)}
              className="flex items-center gap-2 px-6 py-3 bg-slate-800/50 border border-slate-700 rounded-xl hover:bg-slate-800 transition-all"
            >
              <Filter className="w-5 h-5" />
              Filters
            </button>
          </div>

          {/* Filter Options */}
          {showFilters && (
            <div className="grid md:grid-cols-2 gap-4 p-4 bg-slate-800/30 border border-slate-700 rounded-xl">
              <div>
                <label className="block text-slate-300 text-sm font-medium mb-2">
                  Difficulty
                </label>
                <select
                  value={selectedDifficulty}
                  onChange={(e) => setSelectedDifficulty(e.target.value)}
                  className="w-full px-4 py-2 bg-slate-800 border border-slate-700 rounded-lg text-white focus:outline-none focus:border-blue-500"
                >
                  {difficulties.map((diff) => (
                    <option key={diff} value={diff.toLowerCase()}>
                      {diff.charAt(0).toUpperCase() + diff.slice(1)}
                    </option>
                  ))}
                </select>
              </div>

              <div>
                <label className="block text-slate-300 text-sm font-medium mb-2">
                  Category
                </label>
                <select
                  value={selectedCategory}
                  onChange={(e) => setSelectedCategory(e.target.value)}
                  className="w-full px-4 py-2 bg-slate-800 border border-slate-700 rounded-lg text-white focus:outline-none focus:border-blue-500"
                >
                  {categories.map((cat) => (
                    <option key={cat} value={cat.toLowerCase()}>
                      {cat.charAt(0).toUpperCase() + cat.slice(1)}
                    </option>
                  ))}
                </select>
              </div>
            </div>
          )}
        </div>

        {/* Problem List */}
        <div className="space-y-3">
          {filteredProblems.map((item, index) => (
            <div
              onClick={() =>
                (window.location.href = `/problem/${item.problem.Id}`)
              }
              key={item.problem.Id}
              className="group bg-slate-800/30 border border-slate-700 rounded-xl p-5 hover:bg-slate-800/50 hover:border-blue-500/50 transition-all cursor-pointer"
            >
              <div className="flex items-start gap-4">
                {/* Status Icon */}
                <div className="flex-shrink-0 mt-1">
                  {item.isAccept ? (
                    <CheckCircle2 className="w-6 h-6 text-green-400" />
                  ) : (
                    <div className="w-6 h-6 rounded-full border-2 border-slate-600"></div>
                  )}
                </div>

                {/* Problem Info */}
                <div className="flex-1 min-w-0">
                  <div className="flex items-start justify-between gap-4 mb-3">
                    <div className="flex-1">
                      <div className="flex items-center gap-3 mb-2">
                        <span className="text-slate-400 font-mono text-sm">
                          #{item.problem.Id}
                        </span>
                        {item.problem.IsDailyToday && (
                          <span className="flex items-center gap-1 px-2 py-1 bg-yellow-400/10 text-yellow-400 rounded-md text-xs font-medium border border-yellow-400/20">
                            <Star className="w-3 h-3" />
                            Daily +{item.problem.PointDaily}
                          </span>
                        )}
                      </div>
                      <h3 className="text-lg font-semibold text-white group-hover:text-blue-400 transition-colors mb-2">
                        {item.problem.Title}
                      </h3>
                      <div className="flex flex-wrap items-center gap-2">
                        <span
                          className={`px-3 py-1 rounded-full text-xs font-medium border ${getDifficultyColor(
                            item.problem.Difficult
                          )}`}
                        >
                          {item.problem.Difficult}
                        </span>
                        <span className="px-3 py-1 rounded-full text-xs font-medium bg-slate-700/50 text-slate-300 border border-slate-600">
                          {item.problem.Category}
                        </span>
                        <div className="flex items-center gap-1 text-slate-400 text-xs">
                          <Clock className="w-3 h-3" />
                          {item.problem.TimeLimit}ms
                        </div>
                        <div className="flex items-center gap-1 text-slate-400 text-xs">
                          <Zap className="w-3 h-3" />
                          {item.problem.MemoryLimit}MB
                        </div>
                      </div>
                    </div>

                    <ChevronRight className="w-5 h-5 text-slate-500 group-hover:text-blue-400 group-hover:translate-x-1 transition-all flex-shrink-0" />
                  </div>
                </div>
              </div>
            </div>
          ))}
        </div>

        {/* No Results */}
        {filteredProblems.length === 0 && (
          <div className="text-center py-20">
            <div className="w-20 h-20 bg-slate-800/50 rounded-full flex items-center justify-center mx-auto mb-4">
              <Search className="w-10 h-10 text-slate-600" />
            </div>
            <h3 className="text-xl font-semibold text-slate-300 mb-2">
              No problems found
            </h3>
            <p className="text-slate-500">
              Try adjusting your filters or search term
            </p>
          </div>
        )}
      </div>
    </div>
  );
}
