$('.search-box__button').click(function() {
    searchKeyword();
});

$('.search-box input').keypress(function (e) {
    if (e.which == 13) {
        searchKeyword()
    }
});

var url = new URL(window.location.href);
if (url.searchParams.get('keyword') != null ) {
    $('.search-box input').val(url.searchParams.get('keyword'));
    $('.search-box__remove').show();
}

if (url.searchParams.get('tag') != null ) {
    $('.search-box__remove').show();
}

$('.search-box__remove').click(function (e) {
    // window.location.href = '/'
    var url = new URL(window.location.href);
    url.searchParams.delete('keyword');
    url.searchParams.delete('tag');
    window.location.href = url.href
});

function searchKeyword(){
    var url = new URL(window.location.href);
    var valueSearchBox = $('#search-box').val();
    url.searchParams.set('keyword', valueSearchBox);
    window.location.href = url.href

}