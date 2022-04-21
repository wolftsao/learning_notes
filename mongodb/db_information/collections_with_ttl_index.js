// use <database>;

db.getCollectionNames().filter(
    function (c) {
        return db.getCollection(c).getIndexes().filter(
            function (i) {
                return i.hasOwnProperty('expireAfterSeconds');
            }
        ).length > 0
    }
);