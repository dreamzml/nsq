var Pubsub = require('../lib/pubsub');
var AppState = require('../app_state');
var $ = require('jquery');

var BaseView = require('./base');
var RestChannels = require('../collections/rest-channels');
var RestChannel = require('../models/rest-channel');

var RestChannelsView = BaseView.extend({
    className: 'rest-channels container-fluid',

    template: require('./spinner.hbs'),
    events: {
        'click .rest-channel-form button': 'onCreateRestChannel',
        'click .table-condensed button.delete-link': 'onDeleteRestChannel'
    },
    initialize: function() {
        BaseView.prototype.initialize.apply(this, arguments);
        var isAdmin = arguments[0]['isAdmin'];
        this.listenTo(AppState, 'change:graph_interval', this.render);
        this.collection = new RestChannels();
        this.collection.fetch()
            .done(function(data) {
                this.template = require('./rest-channels.hbs');
                this.render({
                    'message': data['message'],
                    'isAdmin': isAdmin
                });
            }.bind(this))
            .fail(this.handleViewError.bind(this))
            .always(Pubsub.trigger.bind(Pubsub, 'view:ready'));
    },
    
    onCreateRestChannel: function(e) {
        debugger;
        e.preventDefault();
        e.stopPropagation();
        var topic = $(e.target.form.elements['topic']).val();
        var channel = $(e.target.form.elements['channel']).val();
        var rest_url = $(e.target.form.elements['rest_url']).val();
        var content_type = e.target.form.elements.content_type.value;
        var method = e.target.form.elements.method.value;
        if (topic === '' || channel === '' || method === '' || rest_url === '') {
            return;
        }
        $.post(AppState.apiPath('/rest-channels'), JSON.stringify({
            'topic': topic,
            'channel': channel,
            'method': method,
            'rest_url': rest_url,
            'content_type': content_type
        }))
        .done(function() { window.location.reload(true); })
        .fail(this.handleAJAXError.bind(this));
    },
    
    onDeleteRestChannel: function(e) {
        e.preventDefault();
        e.stopPropagation();
        console.log("id",$(e.target).data('id'))
        var channel = new RestChannel({
            'id': $(e.target).data('id')
        });
        channel.destroy({'dataType': 'text'})
            .done(function() { window.location.reload(true); })
            .fail(this.handleAJAXError.bind(this));
    }
});

module.exports = RestChannelsView;
