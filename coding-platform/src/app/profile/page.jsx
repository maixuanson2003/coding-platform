"use client";

import React, { useState, useEffect } from "react";
import {
  User,
  Mail,
  Code2,
  Trophy,
  Target,
  TrendingUp,
  Award,
  Calendar,
  Star,
  Zap,
} from "lucide-react";

export default function Profile() {
  const [userData, setUserData] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchProfile = async () => {
      setLoading(true);
      try {
        const userId = localStorage.getItem("user_id");

        const response = await fetch(
          `http://localhost:8080/api/user/${userId}`
        );
        const result = await response.json();

        console.log(result);

        setUserData(result.data);
      } catch (error) {
        console.error("Error fetching profile:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchProfile();
  }, []);

  if (loading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-slate-950 via-slate-900 to-slate-950 flex items-center justify-center">
        <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
      </div>
    );
  }

  const { user, userStats } = userData;
  const totalSolved =
    userStats.total_easy + userStats.total_medium + userStats.total_hard;
  const totalByLanguage =
    userStats.total_cpp +
    userStats.total_java +
    userStats.total_js +
    userStats.total_python;

  const languageStats = [
    {
      name: "C++",
      value: userStats.total_cpp,
      color: "from-purple-500 to-purple-600",
      icon: "⚡",
    },
    {
      name: "JavaScript",
      value: userStats.total_js,
      color: "from-yellow-500 to-yellow-600",
      icon: "🟨",
    },
    {
      name: "Python",
      value: userStats.total_python,
      color: "from-blue-500 to-blue-600",
      icon: "🐍",
    },
    {
      name: "Java",
      value: userStats.total_java,
      color: "from-orange-500 to-orange-600",
      icon: "☕",
    },
  ];

  const difficultyStats = [
    {
      name: "Easy",
      value: userStats.total_easy,
      total: 100,
      color: "bg-green-500",
      textColor: "text-green-400",
    },
    {
      name: "Medium",
      value: userStats.total_medium,
      total: 50,
      color: "bg-yellow-500",
      textColor: "text-yellow-400",
    },
    {
      name: "Hard",
      value: userStats.total_hard,
      total: 30,
      color: "bg-red-500",
      textColor: "text-red-400",
    },
  ];

  const achievements = [
    {
      icon: Trophy,
      title: "First Solve",
      desc: "Solved your first problem",
      unlocked: totalSolved > 0,
    },
    {
      icon: Zap,
      title: "Speed Runner",
      desc: "Solved 10 problems",
      unlocked: totalSolved >= 10,
    },
    {
      icon: Star,
      title: "Daily Warrior",
      desc: "Completed daily challenges",
      unlocked: user.PointDaily > 0,
    },
    {
      icon: Award,
      title: "Polyglot",
      desc: "Used 3+ languages",
      unlocked:
        [
          userStats.total_cpp,
          userStats.total_java,
          userStats.total_js,
          userStats.total_python,
        ].filter((v) => v > 0).length >= 3,
    },
  ];

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-950 via-slate-900 to-slate-950 text-white">
      {/* Animated background */}
      <div className="fixed inset-0 overflow-hidden pointer-events-none">
        <div className="absolute w-96 h-96 bg-blue-500/10 rounded-full blur-3xl -top-48 -left-48 animate-pulse"></div>
        <div className="absolute w-96 h-96 bg-purple-500/10 rounded-full blur-3xl top-1/2 right-0 animate-pulse"></div>
      </div>

      <div className="relative max-w-7xl mx-auto px-6 py-12">
        {/* Header Section */}
        <div className="bg-slate-800/30 border border-slate-700 rounded-2xl p-8 mb-8 backdrop-blur-sm">
          <div className="flex flex-col md:flex-row items-center md:items-start gap-8">
            {/* Avatar */}
            <div className="relative">
              <div className="w-32 h-32 bg-gradient-to-br from-blue-500 to-purple-600 rounded-2xl flex items-center justify-center text-5xl font-bold">
                {user.Avatar === "check" ? (
                  <User className="w-16 h-16 text-white" />
                ) : (
                  user.Username.charAt(0).toUpperCase()
                )}
              </div>
              <div className="absolute -bottom-2 -right-2 w-12 h-12 bg-gradient-to-br from-green-500 to-emerald-600 rounded-lg flex items-center justify-center border-4 border-slate-900">
                <Trophy className="w-6 h-6" />
              </div>
            </div>

            {/* User Info */}
            <div className="flex-1 text-center md:text-left">
              <h1 className="text-4xl font-bold mb-2">{user.Username}</h1>
              <div className="flex flex-col md:flex-row items-center md:items-center gap-4 text-slate-400 mb-4">
                <div className="flex items-center gap-2">
                  <Mail className="w-4 h-4" />
                  {user.Email}
                </div>
                <div className="hidden md:block w-px h-4 bg-slate-600"></div>
                <div className="flex items-center gap-2">
                  <Code2 className="w-4 h-4" />
                  User #{user.Id}
                </div>
              </div>

              {/* Quick Stats */}
              <div className="flex flex-wrap justify-center md:justify-start gap-4">
                <div className="px-4 py-2 bg-slate-700/50 rounded-lg border border-slate-600">
                  <div className="text-2xl font-bold text-blue-400">
                    {totalSolved}
                  </div>
                  <div className="text-xs text-slate-400">Problems Solved</div>
                </div>
                <div className="px-4 py-2 bg-slate-700/50 rounded-lg border border-slate-600">
                  <div className="text-2xl font-bold text-purple-400">
                    {user.NumberHandle}
                  </div>
                  <div className="text-xs text-slate-400">Handles</div>
                </div>
                <div className="px-4 py-2 bg-slate-700/50 rounded-lg border border-slate-600">
                  <div className="text-2xl font-bold text-yellow-400">
                    {user.PointDaily}
                  </div>
                  <div className="text-xs text-slate-400">Daily Points</div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div className="grid lg:grid-cols-2 gap-8 mb-8">
          {/* Language Statistics */}
          <div className="bg-slate-800/30 border border-slate-700 rounded-2xl p-6 backdrop-blur-sm">
            <h2 className="text-2xl font-bold mb-6 flex items-center gap-2">
              <Code2 className="w-6 h-6 text-blue-400" />
              Languages Used
            </h2>
            <div className="space-y-4">
              {languageStats.map((lang, index) => (
                <div key={index}>
                  <div className="flex items-center justify-between mb-2">
                    <div className="flex items-center gap-2">
                      <span className="text-2xl">{lang.icon}</span>
                      <span className="font-semibold">{lang.name}</span>
                    </div>
                    <div className="flex items-center gap-2">
                      <span className="text-2xl font-bold text-blue-400">
                        {lang.value}
                      </span>
                      <span className="text-slate-400 text-sm">
                        (
                        {totalByLanguage > 0
                          ? ((lang.value / totalByLanguage) * 100).toFixed(0)
                          : 0}
                        %)
                      </span>
                    </div>
                  </div>
                  <div className="h-3 bg-slate-700 rounded-full overflow-hidden">
                    <div
                      className={`h-full bg-gradient-to-r ${lang.color} transition-all duration-500`}
                      style={{
                        width: `${
                          totalByLanguage > 0
                            ? (lang.value / totalByLanguage) * 100
                            : 0
                        }%`,
                      }}
                    ></div>
                  </div>
                </div>
              ))}
            </div>
          </div>

          {/* Difficulty Statistics */}
          <div className="bg-slate-800/30 border border-slate-700 rounded-2xl p-6 backdrop-blur-sm">
            <h2 className="text-2xl font-bold mb-6 flex items-center gap-2">
              <Target className="w-6 h-6 text-green-400" />
              Problem Difficulty
            </h2>
            <div className="space-y-6">
              {difficultyStats.map((stat, index) => (
                <div key={index}>
                  <div className="flex items-center justify-between mb-2">
                    <span className="font-semibold">{stat.name}</span>
                    <span className={`text-sm ${stat.textColor}`}>
                      {stat.value} / {stat.total}
                    </span>
                  </div>
                  <div className="h-3 bg-slate-700 rounded-full overflow-hidden">
                    <div
                      className={`h-full ${stat.color} transition-all duration-500`}
                      style={{ width: `${(stat.value / stat.total) * 100}%` }}
                    ></div>
                  </div>
                  <div className="text-right text-xs text-slate-400 mt-1">
                    {((stat.value / stat.total) * 100).toFixed(1)}% Complete
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>

        {/* Progress Overview */}
        <div className="bg-slate-800/30 border border-slate-700 rounded-2xl p-6 mb-8 backdrop-blur-sm">
          <h2 className="text-2xl font-bold mb-6 flex items-center gap-2">
            <TrendingUp className="w-6 h-6 text-purple-400" />
            Overall Progress
          </h2>
          <div className="grid md:grid-cols-3 gap-6">
            <div className="text-center p-6 bg-slate-700/30 rounded-xl border border-slate-600">
              <div className="text-5xl font-bold bg-gradient-to-r from-green-400 to-emerald-400 bg-clip-text text-transparent mb-2">
                {totalSolved}
              </div>
              <div className="text-slate-400 mb-4">Total Solved</div>
              <div className="h-2 bg-slate-700 rounded-full overflow-hidden">
                <div
                  className="h-full bg-gradient-to-r from-green-500 to-emerald-500"
                  style={{ width: "45%" }}
                ></div>
              </div>
            </div>

            <div className="text-center p-6 bg-slate-700/30 rounded-xl border border-slate-600">
              <div className="text-5xl font-bold bg-gradient-to-r from-blue-400 to-cyan-400 bg-clip-text text-transparent mb-2">
                {totalByLanguage}
              </div>
              <div className="text-slate-400 mb-4">Submissions</div>
              <div className="h-2 bg-slate-700 rounded-full overflow-hidden">
                <div
                  className="h-full bg-gradient-to-r from-blue-500 to-cyan-500"
                  style={{ width: "60%" }}
                ></div>
              </div>
            </div>

            <div className="text-center p-6 bg-slate-700/30 rounded-xl border border-slate-600">
              <div className="text-5xl font-bold bg-gradient-to-r from-purple-400 to-pink-400 bg-clip-text text-transparent mb-2">
                {((totalSolved / 180) * 100).toFixed(0)}%
              </div>
              <div className="text-slate-400 mb-4">Completion Rate</div>
              <div className="h-2 bg-slate-700 rounded-full overflow-hidden">
                <div
                  className="h-full bg-gradient-to-r from-purple-500 to-pink-500"
                  style={{ width: `${(totalSolved / 180) * 100}%` }}
                ></div>
              </div>
            </div>
          </div>
        </div>

        {/* Achievements */}
        <div className="bg-slate-800/30 border border-slate-700 rounded-2xl p-6 backdrop-blur-sm">
          <h2 className="text-2xl font-bold mb-6 flex items-center gap-2">
            <Award className="w-6 h-6 text-yellow-400" />
            Achievements
          </h2>
          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-4">
            {achievements.map((achievement, index) => (
              <div
                key={index}
                className={`p-4 rounded-xl border transition-all ${
                  achievement.unlocked
                    ? "bg-gradient-to-br from-yellow-900/30 to-orange-900/30 border-yellow-500/50"
                    : "bg-slate-700/30 border-slate-600 opacity-50"
                }`}
              >
                <div
                  className={`w-12 h-12 rounded-lg flex items-center justify-center mb-3 ${
                    achievement.unlocked
                      ? "bg-gradient-to-br from-yellow-500 to-orange-500"
                      : "bg-slate-600"
                  }`}
                >
                  <achievement.icon className="w-6 h-6 text-white" />
                </div>
                <h3 className="font-semibold mb-1">{achievement.title}</h3>
                <p className="text-sm text-slate-400">{achievement.desc}</p>
                {achievement.unlocked && (
                  <div className="mt-2 text-xs text-yellow-400 flex items-center gap-1">
                    <Star className="w-3 h-3" />
                    Unlocked
                  </div>
                )}
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}
