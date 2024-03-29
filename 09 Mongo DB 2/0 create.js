db.books.insert([
    {_id: 1, title: "Wawasan Pancasila", authorID: 1, publisherID: 1, price: 71200, stats: {page: 324, weight: 300}, publishedAt: new Date("2018-10-01"), category: ["social, politics"]}, 
    {_id: 2, title: "Mata Air Keteladanan", authorID: 1, publisherID: 2, price: 106250, stats: {page: 672, weight: 650}, publishedAt: new Date("2017-09-01"), category: ["social, politics"]},
    {_id: 3, title: "Revolusi Pancasila", authorID: 1, publisherID: 2, price: 54400, stats: {page: 220, weight: 500}, publishedAt: new Date("2015-05-01"), category: ["social, politics"]}, 
    {_id: 4, title: "Self Driving", authorID: 2, publisherID: 1, price: 58650, stats: {page: 286, weight: 300}, publishedAt: new Date("2018-05-01"), category: ["self-development"]},
    {_id: 5, title: "Self Disruption", authorID: 2, publisherID: 2, price: 83300, stats: {page: 400, weight: 800}, publishedAt: new Date("2018-05-01"), category: ["self-development"]} 
]);

db.authors.insert([
    {_id: 1, firstName: "Yudi", lastName: "Latif"},
    {_id: 2, firstName: "Rhenald", lastName: "Kasali"}
]);

db.publishers.insert([
    {_id: 1, publisherName: "Expose"}, {_id: 2, publisherName: "Mizan"}
]);

