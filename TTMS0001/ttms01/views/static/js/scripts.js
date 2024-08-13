
const todayBoxOfficeData = [
    { title: "电影标题1", director: "导演1", image: "movie1.jpg" },
    { title: "电影标题2", director: "导演2", image: "movie2.jpg" },
    { title: "电影标题3", director: "导演3", image: "movie3.jpg" },
    { title: "电影标题4", director: "导演4", image: "movie4.jpg" },
    { title: "电影标题5", director: "导演5", image: "movie5.jpg" }
];

const mostAnticipatedData = [
    { title: "即将上映电影1", director: "导演6", image: "upcoming1.jpg" },
    { title: "即将上映电影2", director: "导演7", image: "upcoming2.jpg" },
    { title: "即将上映电影3", director: "导演8", image: "upcoming3.jpg" },
    { title: "即将上映电影4", director: "导演9", image: "upcoming4.jpg" },
    { title: "即将上映电影5", director: "导演10", image: "upcoming5.jpg" }
];

const topRankedData = [
    { title: "Top电影1", director: "导演11", image: "top1.jpg" },
    { title: "Top电影2", director: "导演12", image: "top2.jpg" },
    { title: "Top电影3", director: "导演13", image: "top3.jpg" },
    { title: "Top电影4", director: "导演14", image: "top4.jpg" },
    { title: "Top电影5", director: "导演15", image: "top5.jpg" }
];

// 生成排行榜项
function generateRankingItems(containerId, data) {
    const container = document.getElementById(containerId);
    container.innerHTML = ""; // 清空现有内容

    data.forEach((item, index) => {
        if (index < 5) { // 只生成前五条信息
            const listItem = document.createElement('li');
            listItem.classList.add('ranking-item');
            listItem.innerHTML = `
                <img src="${item.image}" alt="${item.title}">
                <div class="ranking-item-info">
                    <div class="ranking-item-title">${item.title}</div>
                    <div class="ranking-item-meta">导演：${item.director}</div>
                    <!-- 其他信息 -->
                </div>
            `;
            container.appendChild(listItem);
        }
    });
}

// 在页面加载后生成排行榜数据
document.addEventListener('DOMContentLoaded', () => {
    generateRankingItems('today-box-office', todayBoxOfficeData);
    generateRankingItems('most-anticipated', mostAnticipatedData);
    generateRankingItems('top-ranked', topRankedData);
});

