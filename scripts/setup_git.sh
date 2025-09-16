#!/bin/bash

# Git 远程仓库连接脚本

echo "🚀 Git 远程仓库连接助手"
echo "========================"
echo ""

# 检查是否已经是 Git 仓库
if [ ! -d ".git" ]; then
    echo "❌ 当前目录不是 Git 仓库"
    echo "请先运行: git init"
    exit 1
fi

echo "✅ Git 仓库已初始化"
echo ""

# 显示当前状态
echo "📊 当前 Git 状态:"
git status --short
echo ""

# 检查是否有未提交的更改
if ! git diff --quiet || ! git diff --cached --quiet; then
    echo "⚠️  检测到未提交的更改"
    echo "请先提交所有更改:"
    echo "  git add ."
    echo "  git commit -m 'Your commit message'"
    echo ""
    read -p "是否现在提交所有更改? (y/n): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        git add .
        git commit -m "Update project files"
        echo "✅ 更改已提交"
    else
        echo "❌ 请先提交更改后再运行此脚本"
        exit 1
    fi
fi

echo ""

# 显示当前远程仓库
echo "🔗 当前远程仓库:"
git remote -v
echo ""

# 询问用户选择
echo "请选择要连接的远程仓库平台:"
echo "1) GitHub"
echo "2) GitLab"
echo "3) Gitee (码云)"
echo "4) 其他 (手动输入 URL)"
echo "5) 查看帮助"
echo ""

read -p "请输入选择 (1-5): " choice

case $choice in
    1)
        echo ""
        echo "📝 GitHub 设置"
        echo "============="
        echo "1. 访问 https://github.com"
        echo "2. 点击 'New repository'"
        echo "3. 填写仓库信息"
        echo "4. 不要勾选任何初始化选项"
        echo "5. 点击 'Create repository'"
        echo ""
        read -p "请输入您的 GitHub 用户名: " username
        read -p "请输入仓库名称: " repo_name
        
        remote_url="https://github.com/$username/$repo_name.git"
        echo ""
        echo "将添加远程仓库: $remote_url"
        ;;
    2)
        echo ""
        echo "📝 GitLab 设置"
        echo "============="
        read -p "请输入您的 GitLab 用户名: " username
        read -p "请输入仓库名称: " repo_name
        
        remote_url="https://gitlab.com/$username/$repo_name.git"
        echo ""
        echo "将添加远程仓库: $remote_url"
        ;;
    3)
        echo ""
        echo "📝 Gitee 设置"
        echo "============"
        read -p "请输入您的 Gitee 用户名: " username
        read -p "请输入仓库名称: " repo_name
        
        remote_url="https://gitee.com/$username/$repo_name.git"
        echo ""
        echo "将添加远程仓库: $remote_url"
        ;;
    4)
        echo ""
        read -p "请输入完整的仓库 URL: " remote_url
        ;;
    5)
        echo ""
        echo "📚 帮助信息"
        echo "=========="
        echo "请查看 GIT_SETUP_GUIDE.md 文件获取详细说明"
        exit 0
        ;;
    *)
        echo "❌ 无效选择"
        exit 1
        ;;
esac

echo ""

# 添加远程仓库
echo "🔗 添加远程仓库..."
git remote add origin "$remote_url"

if [ $? -eq 0 ]; then
    echo "✅ 远程仓库添加成功"
else
    echo "❌ 添加远程仓库失败"
    echo "可能的原因:"
    echo "  - 仓库 URL 不正确"
    echo "  - 网络连接问题"
    echo "  - 仓库已存在"
    exit 1
fi

echo ""

# 推送到远程仓库
echo "📤 推送代码到远程仓库..."
git push -u origin main

if [ $? -eq 0 ]; then
    echo ""
    echo "🎉 成功！您的项目已连接到远程仓库"
    echo ""
    echo "📋 后续操作:"
    echo "  - 查看仓库: 在浏览器中打开 $remote_url"
    echo "  - 推送更改: git push"
    echo "  - 拉取更改: git pull"
    echo "  - 查看状态: git status"
    echo ""
    echo "🔗 远程仓库信息:"
    git remote -v
else
    echo ""
    echo "❌ 推送失败"
    echo ""
    echo "可能的原因:"
    echo "  - 认证失败 (需要设置 SSH 密钥或 Personal Access Token)"
    echo "  - 网络连接问题"
    echo "  - 仓库权限问题"
    echo ""
    echo "解决方案:"
    echo "  1. 设置 SSH 密钥: https://docs.github.com/en/authentication/connecting-to-github-with-ssh"
    echo "  2. 使用 Personal Access Token: https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token"
    echo "  3. 检查网络连接"
    echo ""
    echo "手动推送命令:"
    echo "  git push -u origin main"
fi
