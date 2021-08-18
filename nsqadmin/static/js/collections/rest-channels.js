var _ = require('underscore');
var Backbone = require('backbone');

var AppState = require('../app_state');

var RestChannel = require('../models/rest-channel');

var RestChannels = Backbone.Collection.extend({
    model: RestChannel,

    comparator: 'id',

    constructor: function RestChannels() {
        Backbone.Collection.prototype.constructor.apply(this, arguments);
    },

    url: function() {
        return AppState.apiPath('/rest-channels');
    },

    parse: function(resp) {
        var RestChannels = _.map(resp['restChannels'], function(item) {
            var createTime = new Date(item.created_at).format("yyyy-MM-dd hh:mm")
            return {
                "id": item.ID,
                "topic": item.topic, 
                "channel":item.channel, 
                "method":item.method, 
                "rest_url":item.rest_url, 
                'content_type': item.content_type,
                "create_at":createTime};
        });
        return RestChannels;
    }
});

module.exports = RestChannels;
