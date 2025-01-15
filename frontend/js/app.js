const URL = `http://localhost:8080/news`
const ERROR_NOTICIAS = 'Error al obtener las noticias:'

fetch(URL)
    .then(response => response.json())
    .then(data => {
        displayNews(data);
    })
    .catch(error => console.error(ERROR_NOTICIAS, error));

function displayNews(data) {
    const newsContainer = document.getElementById('newsContainer');
    for (let i = 0; i < 10 && i < data.articles.length; i++) {
        const article = data.articles[i];
        const newsItem = document.createElement('div');
        newsItem.classList.add('news-item');
        newsItem.innerHTML = `
            <h3>${article.title}</h3>
            <p>${article.description ? article.description : 'No description available.'}</p>
        `;
        newsContainer.appendChild(newsItem);
    }
}
