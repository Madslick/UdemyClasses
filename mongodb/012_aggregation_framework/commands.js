let db;

db.persons.aggregate([
    // Stage 1 matching
    {
        $match: {
            gender: "female"
        }
    },
    // Stage 2 Grouping by a field or certain fields
    {
        $group: {
            _id: { state: "$location.state"}, totalPersons: { $sum: 1 } // _id is a special way to group the fields. apparently it's common in grouping stage
        }
    }
]).pretty()
