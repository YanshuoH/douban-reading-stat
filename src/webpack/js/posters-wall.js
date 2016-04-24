$(function() {
    var template = '<img class="book-cover" src="$imgSrc$" alt="$bookTitle$">'

    var container = $('#posters-wall');
    var posters = dataset['result']['2015']['posters'];
    for (var i = 0; i < posters.length; i ++) {
        var poster = posters[i];
        var elem = template.replace('$imgSrc$', poster['link']);
        elem = elem.replace('$bookTitle$', poster['title']);
        container.append(elem);
    }
});