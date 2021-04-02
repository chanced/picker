# picker

Utilities that pair with the official Elasticsearch Go package

## Todo

Testing is incredibly light right now.

### [Aggregations](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations.html)

An aggregation summarizes your data as metrics, statistics, or other analytics.

- #### [Bucket aggregations](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket.html)

  Bucket aggregations don’t calculate metrics over fields like the metrics aggregations do, but instead, they create buckets of documents. Each bucket is associated with a criterion (depending on the aggregation type) which determines whether or not a document in the current context "falls" into it. In other words, the buckets effectively define document sets. In addition to the buckets themselves, the bucket aggregations also compute and return the number of documents that "fell into" each bucket.

  Bucket aggregations, as opposed to metrics aggregations, can hold sub-aggregations. These sub-aggregations will be aggregated for the buckets created by their "parent" bucket aggregation.

  There are different bucket aggregators, each with a different "bucketing" strategy. Some define a single bucket, some define fixed number of multiple buckets, and others dynamically create the buckets during the aggregation process.

  - [ ] **[Adjacency matrix](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-adjacency-matrix-aggregation.html)**\
         A bucket aggregation returning a form of adjacency matrix. The request provides a collection of named filter expressions, similar to the filters aggregation request. Each bucket in the response represents a non-empty cell in the matrix of intersecting filters.
  - [ ] **[Auto-interval date histogram](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-autodatehistogram-aggregation.html)**\
         A multi-bucket aggregation similar to the Date histogram except instead of providing an interval to use as the width of each bucket, a target number of buckets is provided indicating the number of buckets needed and the interval of the buckets is automatically chosen to best achieve that target. The number of buckets returned will always be less than or equal to this target number.
        The buckets field is optional, and will default to 10 buckets if not specified.
        Requesting a target of 10 buckets.
  - [ ] **[Children](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-children-aggregation.html)**\
         A special single bucket aggregation that selects child documents that have the specified type, as defined in a join field.
  - [ ] **[Composite](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-composite-aggregation.html)**\
         A multi-bucket aggregation that creates composite buckets from different sources.
        Unlike the other multi-bucket aggregations, you can use the composite aggregation to paginate all buckets from a multi-level aggregation efficiently. This aggregation provides a way to stream all buckets of a specific aggregation, similar to what scroll does for documents.
        The composite buckets are built from the combinations of the values extracted/created for each document and each combination is considered as a composite bucket.
  - [ ] **[Date histogram](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-datehistogram-aggregation.html)**\
         This multi-bucket aggregation is similar to the normal histogram, but it can only be used with date or date range values. Because dates are represented internally in Elasticsearch as long values, it is possible, but not as accurate, to use the normal histogram on dates as well. The main difference in the two APIs is that here the interval can be specified using date/time expressions. Time-based data requires special support because time-based intervals are not always a fixed length.
  - [ ] **[Date range](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-daterange-aggregation.html)**\
         A range aggregation that is dedicated for date values. The main difference between this aggregation and the normal range aggregation is that the from and to values can be expressed in Date Math expressions, and it is also possible to specify a date format by which the from and to response fields will be returned. Note that this aggregation includes the from value and excludes the to value for each range.
  - [ ] **[Diversified sampler](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-diversified-sampler-aggregation.html)**\
         Like the sampler aggregation this is a filtering aggregation used to limit any sub aggregations' processing to a sample of the top-scoring documents. The diversified_sampler aggregation adds the ability to limit the number of matches that share a common value such as an "author".
  - [ ] **[Filter](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-filter-aggregation.html)**\
         Defines a single bucket of all the documents in the current document set context that match a specified filter. Often this will be used to narrow down the current aggregation context to a specific set of documents.
  - [ ] **[Filters](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-filters-aggregation.html)**\
         Defines a multi bucket aggregation where each bucket is associated with a filter. Each bucket will collect all documents that match its associated filter.
  - [ ] **[Geo-distance](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-geodistance-aggregation.html)**\
         A multi-bucket aggregation that works on geo_point fields and conceptually works very similar to the range aggregation. The user can define a point of origin and a set of distance range buckets. The aggregation evaluate the distance of each document value from the origin point and determines the buckets it belongs to based on the ranges (a document belongs to a bucket if the distance between the document and the origin falls within the distance range of the bucket).
  - [ ] **[Geohash grid](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-geohashgrid-aggregation.html)**\
         A multi-bucket aggregation that groups geo_point and geo_shape values into buckets that represent a grid. The resulting grid can be sparse and only contains cells that have matching data. Each cell is labeled using a geohash which is of user-definable precision.
  - [ ] **[Geotile grid](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-geotilegrid-aggregation.html)**\
         A multi-bucket aggregation that groups geo_point and geo_shape values into buckets that represent a grid. The resulting grid can be sparse and only contains cells that have matching data. Each cell corresponds to a map tile as used by many online map sites. Each cell is labeled using a "{zoom}/{x}/{y}" format, where zoom is equal to the user-specified precision.
  - [ ] **[Global](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-global-aggregation.html)**\
         Defines a single bucket of all the documents within the search execution context. This context is defined by the indices and the document types you’re searching on, but is not influenced by the search query itself.
  - [ ] **[Histogram](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-histogram-aggregation.html)**\
         A multi-bucket values source based aggregation that can be applied on numeric values or numeric range values extracted from the documents. It dynamically builds fixed size (a.k.a. interval) buckets over the values.
  - [ ] **[IP range](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-iprange-aggregation.html)**\
         Just like the dedicated date range aggregation, there is also a dedicated range aggregation for IP typed fields
  - [ ] **[Missing](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-missing-aggregation.html)**\
         A field data based single bucket aggregation, that creates a bucket of all documents in the current document set context that are missing a field value (effectively, missing a field or having the configured NULL value set). This aggregator will often be used in conjunction with other field data bucket aggregators (such as ranges) to return information for all the documents that could not be placed in any of the other buckets due to missing field data values.
  - [ ] **[Multi Terms](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-multi-terms-aggregation.html) [X-Pack]**\
         A multi-bucket value source based aggregation where buckets are dynamically built - one per unique set of values. The multi terms aggregation is very similar to the terms aggregation, however in most cases it will be slower than the terms aggregation and will consume more memory. Therefore, if the same set of fields is constantly used, it would be more efficient to index a combined key for this fields as a separate field and use the terms aggregation on this field.
  - [ ] **[Nested](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-nested-aggregation.html)**\
         A special single bucket aggregation that enables aggregating nested documents.
  - [ ] **[Parent](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-parent-aggregation.html)**\
         A special single bucket aggregation that selects parent documents that have the specified type, as defined in a join field.
  - [ ] **[Range](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-range-aggregation.html)**\
         A multi-bucket value source based aggregation that enables the user to define a set of ranges - each representing a bucket. During the aggregation process, the values extracted from each document will be checked against each bucket range and "bucket" the relevant/matching document. Note that this aggregation includes the from value and excludes the to value for each range.
  - [ ] **[Rare terms](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-rare-terms-aggregation.html)**\
         A multi-bucket value source based aggregation which finds "rare" terms — terms that are at the long-tail of the distribution and are not frequent. Conceptually, this is like a terms aggregation that is sorted by \_count ascending. As noted in the terms aggregation docs, actually ordering a terms agg by count ascending has unbounded error. Instead, you should use the rare_terms aggregation
  - [ ] **[Reverse nested](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-reverse-nested-aggregation.html)**\
         A special single bucket aggregation that enables aggregating on parent docs from nested documents. Effectively this aggregation can break out of the nested block structure and link to other nested structures or the root document, which allows nesting other aggregations that aren’t part of the nested object in a nested aggregation.
  - [ ] **[Sampler](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-sampler-aggregation.html)**\
         A filtering aggregation used to limit any sub aggregations' processing to a sample of the top-scoring documents.
  - [ ] **[Significant terms](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-significantterms-aggregation.html)**\
         An aggregation that returns interesting or unusual occurrences of terms in a set.
  - [ ] **[Significant text](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-significanttext-aggregation.html)**\
         An aggregation that returns interesting or unusual occurrences of free-text terms in a set.
  - [ ] **[Terms](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-terms-aggregation.html)**\
         A multi-bucket value source based aggregation where buckets are dynamically built - one per unique value.
  - [ ] **[Variable width histogram](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-variablewidthhistogram-aggregation.html)**\
         This is a multi-bucket aggregation similar to Histogram. However, the width of each bucket is not specified. Rather, a target number of buckets is provided and bucket intervals are dynamically determined based on the document distribution. This is done using a simple one-pass document clustering algorithm that aims to obtain low distances between bucket centroids. Unlike other multi-bucket aggregations, the intervals will not necessarily have a uniform width.
  - [ ] **[Subtleties of bucketing range fields](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-range-field-note.html)**\
         Since a range represents multiple values, running a bucket aggregation over a range field can result in the same document landing in multiple buckets. This can lead to surprising behavior, such as the sum of bucket counts being higher than the number of matched documents.

- #### [Metric aggreations](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics.html)

  The aggregations in this family compute metrics based on values extracted in one way or another from the documents that are being aggregated. The values are typically extracted from the fields of the document (using the field data), but can also be generated using scripts.

  Numeric metrics aggregations are a special type of metrics aggregation which output numeric values. Some aggregations output a single numeric metric (e.g. avg) and are called single-value numeric metrics aggregation, others generate multiple metrics (e.g. stats) and are called multi-value numeric metrics aggregation. The distinction between single-value and multi-value numeric metrics aggregations plays a role when these aggregations serve as direct sub-aggregations of some bucket aggregations (some bucket aggregations enable you to sort the returned buckets based on the numeric metrics in each bucket).

  - [ ] **[Avg](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-avg-aggregation.html)**\
         A single-value metrics aggregation that computes the average of numeric values that are extracted from the aggregated documents. These values can be extracted either from specific numeric fields in the documents, or be generated by a provided script.
  - [ ] **[Boxplot](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-boxplot-aggregation.html) [X-Pack]**\
         A boxplot metrics aggregation that computes boxplot of numeric values extracted from the aggregated documents. These values can be generated by a provided script or extracted from specific numeric or histogram fields in the documents.
  - [ ] **[Cardinality](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-cardinality-aggregation.html)**\
         A single-value metrics aggregation that calculates an approximate count of distinct values. Values can be extracted either from specific fields in the document or generated by a script.
  - [ ] **[Extended stats](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-extendedstats-aggregation.html)**\
         A multi-value metrics aggregation that computes stats over numeric values extracted from the aggregated documents. These values can be extracted either from specific numeric fields in the documents, or be generated by a provided script.\
         The extended_stats aggregations is an extended version of the stats aggregation, where additional metrics are added such as sum_of_squares, variance, std_deviation and std_deviation_bounds.
  - [ ] **[Geo-bounds](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-geobounds-aggregation.html)**\
         A metric aggregation that computes the bounding box containing all geo values for a field.
  - [ ] **[Geo-centroid](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-geocentroid-aggregation.html)**\
         metric aggregation that computes the weighted centroid from all coordinate values for geo fields.
  - [ ] **[Geo-Line](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-geo-line.html) [X-Pack]**\
         The geo_line aggregation aggregates all geo_point values within a bucket into a LineString ordered by the chosen sort field. This sort can be a date field, for example. The bucket returned is a valid GeoJSON Feature representing the line geometry.
  - [ ] **[Matrix stats](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-matrix-stats-aggregation.html)**
  - [ ] **[Max](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-max-aggregation.html)**\
         A single-value metrics aggregation that keeps track and returns the maximum value among the numeric values extracted from the aggregated documents. These values can be extracted either from specific numeric fields in the documents, or be generated by a provided script.
  - [ ] **[Median absolute deviation](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-median-absolute-deviation-aggregation.html)**\
         This single-value aggregation approximates the median absolute deviation of its search results.\
         Median absolute deviation is a measure of variability. It is a robust statistic, meaning that it is useful for describing data that may have outliers, or may not be normally distributed. For such data it can be more descriptive than standard deviation.
  - [ ] **[Min](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-min-aggregation.html)**\
         single-value metrics aggregation that keeps track and returns the minimum value among numeric values extracted from the aggregated documents. These values can be extracted either from specific numeric fields in the documents, or be generated by a provided script.
  - [ ] **[Percentile ranks](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-percentile-rank-aggregation.html)**\
         A multi-value metrics aggregation that calculates one or more percentile ranks over numeric values extracted from the aggregated documents. These values can be generated by a provided script or extracted from specific numeric or histogram fields in the documents.
  - [ ] **[Percentiles](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-percentile-aggregation.html)**\
         A multi-value metrics aggregation that calculates one or more percentiles over numeric values extracted from the aggregated documents. These values can be generated by a provided script or extracted from specific numeric or histogram fields in the documents.
  - [ ] **[Rate](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-rate-aggregation.html) [X-Pack]**\
         A rate metrics aggregation can be used only inside a date_histogram and calculates a rate of documents or a field in each date_histogram bucket. The field values can be generated by a provided script or extracted from specific numeric or histogram fields in the documents.
  - [ ] **[Scripted metric](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-scripted-metric-aggregation.html)**\
         A metric aggregation that executes using scripts to provide a metric output.
  - [ ] **[Stats](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-stats-aggregation.html)**\
         A multi-value metrics aggregation that computes stats over numeric values extracted from the aggregated documents. These values can be extracted either from specific numeric fields in the documents, or be generated by a provided script.
        The stats that are returned consist of: min, max, sum, count and avg.
  - [ ] **[String stats](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-string-stats-aggregation.html) [X-Pack]**\
         A multi-value metrics aggregation that computes statistics over string values extracted from the aggregated documents. These values can be retrieved either from specific keyword fields in the documents or can be generated by a provided script.
  - [ ] **[Sum]()**\
  - [ ] **[T-test]()**\
  - [ ] **[Top hits]()**\
  - [ ] **[Top metrics]()**\
  - [ ] **[Value count]()**\
  - [ ] **[Weighted avg]()**\

- #### [Pipeline aggregations](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-pipeline.html)

  Pipeline aggregations work on the outputs produced from other aggregations rather than from document sets, adding information to the output tree. There are many different types of pipeline aggregation, each computing different information from other aggregations, but these types can be broken down into two families:

  **Parent**
  A family of pipeline aggregations that is provided with the output of its parent aggregation and is able to compute new buckets or new aggregations to add to existing buckets.

  **Sibling**
  Pipeline aggregations that are provided with the output of a sibling aggregation and are able to compute a new aggregation which will be at the same level as the sibling aggregation.

  - [ ] **[Average bucket]()**\
  - [ ] **[Bucket script]()**\
  - [ ] **[Bucket selector]()**\
  - [ ] **[Bucket sort]()**\
  - [ ] **[Cumulative cardinality]()**\
  - [ ] **[Cumulative sum]()**\
  - [ ] **[Derivative]()**\
  - [ ] **[Extended stats bucket]()**\
  - [ ] **[Inference bucket]()**\
  - [ ] **[Max bucket]()**\
  - [ ] **[Min bucket]()**\
  - [ ] **[Moving average]()**\
  - [ ] **[Moving function]()**\
  - [ ] **[Moving percentiles]()**\
  - [ ] **[Normalize]()**\
  - [ ] **[Percentiles bucket]()**\
  - [ ] **[Serial differencing]()**\
  - [ ] **[Stats bucket]()**\
  - [ ] **[Sum bucket]()**\

### [Analyzers](https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-analyzers.html)

- #### Built-in analyzers
  - [ ] **[Standard](https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-standard-analyzer.html)**\
         The standard analyzer divides text into terms on word boundaries, as defined by the Unicode Text Segmentation algorithm. It removes most punctuation, lowercases terms, and supports removing stop words.
  - [ ] **[Simple](https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-simple-analyzer.html)**\
         The simple analyzer divides text into terms whenever it encounters a character which is not a letter. It lowercases all terms.
  - [ ] **[Whitespace](https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-whitespace-analyzer.html)**\
         The whitespace analyzer divides text into terms whenever it encounters any whitespace character. It does not lowercase terms.
  - [ ] **[Stop](https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-stop-analyzer.html)**\
         The stop analyzer is like the simple analyzer, but also supports removal of stop words.
  - [ ] **[Keyword](https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-keyword-analyzer.html)**\
         The keyword analyzer is a “noop” analyzer that accepts whatever text it is given and outputs the exact same text as a single term.
  - [ ] **[Pattern](https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-pattern-analyzer.html)**\
         The pattern analyzer uses a regular expression to split the text into terms. It supports lower-casing and stop words.
  - [ ] **[Languages](https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-lang-analyzer.html)**\
         Elasticsearch provides many language-specific analyzers like english or french.
  - [ ] **[Fingerprint](https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-fingerprint-analyzer.html)**\
         The fingerprint analyzer is a specialist analyzer which creates a fingerprint which can be used for duplicate detection.

### [Field Mappings](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/mapping-types.html)

- #### Common types

  - [ ] **[Binary](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/binary.html)**\
         Binary value encoded as a Base64 string.
  - [ ] **[Boolean](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/boolean.html)**\
         true and false values.
  - [ ] **[Keyword](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/keyword.html#keyword-field-type)**\
         used for structured content such as IDs, email addresses, hostnames, status codes, zip codes, or tags.
  - [ ] **[Constant keyword](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/keyword.html#constant-keyword-field-type) [X-Pack]**\
         Constant keyword is a specialization of the keyword field for the case that all documents in the index have the same value.
  - [ ] **[Wildcard](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/keyword.html#wildcard-field-type) [X-Pack]**\
         The wildcard field type is a specialized keyword field for unstructured machine-generated content you plan to search using grep-like wildcard and regexp queries. The wildcard type is optimized for fields with large values or high cardinality.
  - [ ] **[Numbers](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/number.html)**\
         Numeric types, such as long and double, used to express amounts.
  - [ ] **[Date](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/date.html)**\
         Date field type
  - [ ] **[Date nanoseconds](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/date_nanos.html)**\
         Date nanoseconds field type
  - [ ] **[Alias](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/alias.html)**\
         Defines an alias for an existing field.

- #### Objects and relational types

  - [ ] **[Object](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/object.html)**\
         A JSON object.
  - [ ] **[Flattened](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/flattened.html)**\
         An entire JSON object as a single field value.
  - [ ] **[Nested](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/nested.html)**\
         A JSON object that preserves the relationship between its subfields.
  - [ ] **[Join](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/parent-join.html)**\
         Defines a parent/child relationship for documents in the same index.

- #### Structured data types

  - [ ] **[Range](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/range.html)**\
         Range types, such as long_range, double_range, date_range, and ip_range.
  - [ ] **[IP](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/ip.html)**\
         IPv4 and IPv6 addresses.
  - [ ] **[Version](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/version.html) [X-Pack]**\
         Software versions. Supports Semantic Versioning precedence rules.
  - [ ] **[Murmur3](https://www.elastic.co/guide/en/elasticsearch/plugins/7.12/mapper-murmur3.html) [X-Pack]**\
         Compute and stores hashes of values.

- #### Aggregate data types

  - [ ] **[Aggregate metric double](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/aggregate-metric-double.html) [X-Pack]**\
         Pre-aggregated metric values.
  - [ ] **[Histogram](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/histogram.html) [X-Pack]**\
         Pre-aggregated numerical values in the form of a histogram.

- #### Text search types

  - [ ] **[Text](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/text.html)**\
         Analyzed, unstructured text.
  - [ ] **[Annotated-text](https://www.elastic.co/guide/en/elasticsearch/plugins/7.12/mapper-annotated-text.html)**\
         Text containing special markup. Used for identifying named entities.
  - [ ] **[Completion](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/search-suggesters.html#completion-suggester)**\
         Used for auto-complete suggestions.
  - [ ] **[Search as you type](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/search-as-you-type.html)**\
         text-like type for as-you-type completion.
  - [ ] **[Token count](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/token-count.html)**\
         A count of tokens in a text.

- #### Document ranking types

  - [ ] **[Dense vector](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/dense-vector.html) [X-Pack]**\
         Records dense vectors of float values.
  - [ ] **[Sparse vector](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/sparse-vector.html) [X-Pack] [Deprecated]**\
         Records sparse vectors of float values.
  - [ ] **[Rank feature](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/rank-feature.html)**\
         Records a numeric feature to boost hits at query time.
  - [ ] **[Rank features](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/rank-features.html)**\
         Records numeric features to boost hits at query time.

- #### Spatial data types

  - [ ] **[Geo point](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/geo-point.html)**\
         Latitude and longitude points.
  - [ ] **[Geo shape](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/geo-shape.html)**\
         Complex shapes, such as polygons.
  - [ ] **[Point](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/point.html)**\
         Arbitrary cartesian points.
  - [ ] **[Shape](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/shape.html)**\
         Arbitrary cartesian geometries.

- #### Other types
  - [ ] **[Percolator](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/percolator.html)**\
         Indexes queries written in Query DSL.

### Queries

- #### [Compound queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/compound-queries.html)
  Compound queries wrap other compound or leaf queries, either to combine their results and scores, to change their behaviour, or to switch from query to filter context.
  - [x] **[Boolean](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-bool-query.html)**\
         The default query for combining multiple leaf or compound query clauses, as must, should, must_not, or filter clauses. The must and should clauses have their scores combined — the more matching clauses, the better — while the must_not and filter clauses are executed in filter context.
  - [ ] **[Boosting](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-boosting-query.html)**\
         Return documents which match a positive query, but reduce the score of documents which also match a negative query.
  - [ ] **[Constant score](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-constant-score-query.html)**\
         A query which wraps another query, but executes it in filter context. All matching documents are given the same “constant” \_score.
  - [ ] **[Disjunction max](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-dis-max-query.html)**\
         A query which accepts multiple queries, and returns any documents which match any of the query clauses. While the bool query combines the scores from all matching queries, the dis_max query uses the score of the single best- matching query clause.
  - [x] **[Function score](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-function-score-query.html)**\
         Modify the scores returned by the main query with functions to take into account factors like popularity, recency, distance, or custom algorithms implemented with scripting.
- #### [Fulltext queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/full-text-queries.html)
  The full text queries enable you to search analyzed text fields such as the body of an email. The query string is processed using the same analyzer that was applied to the field during indexing.
  - [ ] **[Intervals](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-intervals-query.html)**\
         A full text query that allows fine-grained control of the ordering and proximity of matching terms.
  - [x] **[Match](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-match-query.html)**\
         The standard query for performing full text queries, including fuzzy matching and phrase or proximity queries.
  - [ ] **[Match bool prefix](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-match-bool-prefix-query.html)**\
         Creates a bool query that matches each term as a term query, except for the last term, which is matched as a prefix query
  - [ ] **[Match phrase](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-match-query-phrase.html)**\
         Like the match query but used for matching exact phrases or word proximity matches.
  - [ ] **[Match phrase prefix](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-match-query-phrase-prefix.html)**\
         Like the match_phrase query, but does a wildcard search on the final word.
  - [ ] **[Multi-match](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-multi-match-query.html)**\
         The multi-field version of the match query.
  - [ ] **[Common Terms](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-common-terms-query.html) [Deprecated]**\
         A more specialized query which gives more preference to uncommon words.
  - [ ] **[Query string](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-query-string-query.html)**\
         Supports the compact Lucene query string syntax, allowing you to specify AND|OR|NOT conditions and multi-field search within a single query string. For expert users only.
  - [ ] **[Simple query string](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-simple-query-string-query.html)**\
         A simpler, more robust version of the query_string syntax suitable for exposing directly to users.
- #### [Geo queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/geo-queries.html)
  Elasticsearch supports two types of geo data: geo_point fields which support lat/lon pairs, and geo_shape fields, which support points, lines, circles, polygons, multi-polygons, etc.
  - [ ] **[Geo bounding box](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-geo-bounding-box-query.html)**\
         Finds documents with geo-points that fall into the specified rectangle.
  - [ ] **[Geo distance](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-geo-distance-query.html)**\
         Finds documents with geo-points within the specified distance of a central point.
  - [ ] **[Geo polygon](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-geo-polygon-query.html) [Deprecated]**\
         Find documents with geo-points within the specified polygon.
  - [ ] **[Geo shape](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-geo-shape-query.html)**
    - geo-shapes which either intersect, are contained by, or do not intersect with the specified geo-shape
    - geo-points which intersect the specified geo-shape
- #### [Shape queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/shape-queries.html) [X-Pack]
  Like geo_shape Elasticsearch supports the ability to index arbitrary two dimension (non Geospatial) geometries making it possible to map out virtual worlds, sporting venues, theme parks, and CAD diagrams.
  Elasticsearch supports two types of cartesian data: point fields which support x/y pairs, and shape fields, which support points, lines, circles, polygons, multi-polygons, etc.
  - [ ] **[Shape](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-shape-query.html)**
    - shapes which either intersect, are contained by, are within or do not intersect with the specified shape
    - points which intersect the specified shape
- #### [Joining queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/joining-queries.html)
  - [ ] **[Nested](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-nested-query.html)**\
         Documents may contain fields of type nested. These fields are used to index arrays of objects, where each object can be queried (with the nested query) as an independent document.
  - [ ] **[Has child](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-has-child-query.html)**\
         A join field relationship can exist between documents within a single index. The has_child query returns parent documents whose child documents match the specified query, while the has_parent query returns child documents whose parent document matches the specified query.
  - [ ] **[Has parent](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-has-parent-query.html)**\
         Returns child documents whose joined parent document matches a provided query. You can create parent-child relationships between documents in the same index using a join field mapping.
  - [ ] **[Parent ID](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-parent-id-query.html)**\
         Returns child documents joined to a specific parent document. You can use a join field mapping to create parent-child relationships between documents in the same index.
- [x] **[Match all](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-match-all-query.html)**\
       The most simple query, which matches all documents, giving them all a \_score of 1.0.
- [x] **[Match none](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-match-all-query.html)**\
       This is the inverse of the match_all query, which matches no documents.
- #### [Span queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/span-queries.html)

  Span queries are low-level positional queries which provide expert control over the order and proximity of the specified terms. These are typically used to implement very specific queries on legal documents or patents.

  It is only allowed to set boost on an outer span query. Compound span queries, like span_near, only use the list of matching spans of inner span queries in order to find their own spans, which they then use to produce a score. Scores are never computed on inner span queries, which is the reason why boosts are not allowed: they only influence the way scores are computed, not spans.

  Span queries cannot be mixed with non-span queries (with the exception of the span_multi query).

  - [ ] **[Span containing](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-containing-query.html)**\
         Accepts a list of span queries, but only returns those spans which also match a second span query.
  - [ ] **[Field masking span](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-field-masking-query.html)**\
         Allows queries like span-near or span-or across different fields.
  - [ ] **[Span first](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-first-query.html)**\
         Accepts another span query whose matches must appear within the first N positions of the field.
  - [ ] **[Span multi](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-multi-term-query.html)**\
         Wraps a term, range, prefix, wildcard, regexp, or fuzzy query.
  - [ ] **[Span near](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-near-query.html)**\
         Accepts multiple span queries whose matches must be within the specified distance of each other, and possibly in the same order.
  - [ ] **[Span not](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-not-query.html)**\
         Wraps another span query, and excludes any documents which match that query.
  - [ ] **[Span or](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-or-query.html)**\
         Combines multiple span queries — returns documents which match any of the specified queries.
  - [ ] **[Span term](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-term-query.html)**\
         The equivalent of the term query but for use with other span queries.
  - [ ] **[Span within](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-within-query.html)**\
         The result from a single span query is returned as long is its span falls within the spans returned by a list of other span queries.

- #### [Specialized queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/specialized-queries.html)
  - [ ] **[Distance feature](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-distance-feature-query.html)**\
         A query that computes scores based on the dynamically computed distances between the origin and documents' date, date_nanos and geo_point fields. It is able to efficiently skip non-competitive hits.
  - [ ] **[More like this](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-mlt-query.html)**\
         This query finds documents which are similar to the specified text, document, or collection of documents.
  - [ ] **[Percolate](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-percolate-query.html)**\
         This query finds queries that are stored as documents that match with the specified document.
  - [ ] **[Rank feature](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-rank-feature-query.html)**\
         A query that computes scores based on the values of numeric features and is able to efficiently skip non-competitive hits.
  - [x] **[Script](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-script-query.html)**\
         This query allows a script to act as a filter. Also see the function_score query.
  - [x] **[Script score](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-script-score-query.html)**\
         A query that allows to modify the score of a sub-query with a script.
  - [ ] **[Wrapper](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-wrapper-query.html)**\
         A query that accepts other queries as json or yaml string.
  - [ ] **[Pinned](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-pinned-query.html)**\
         A query that promotes selected documents over others matching a given query.
- #### [Term-level queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/term-level-queries.html)
  You can use term-level queries to find documents based on precise values in structured data. Examples of structured data include date ranges, IP addresses, prices, or product IDs.
  Unlike full-text queries, term-level queries do not analyze search terms. Instead, term-level queries match the exact terms stored in a field.
  - [x] **[Exists](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-exists-query.html)**\
         Returns documents that contain any indexed value for a field.
  - [x] **[Fuzzy](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-fuzzy-query.html)**\
         Returns documents that contain terms similar to the search term. Elasticsearch measures similarity, or fuzziness, using a Levenshtein edit distance.
  - [ ] **[IDs](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-ids-query.html)**\
         Returns documents based on their document IDs.
  - [x] **[Prefix](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-prefix-query.html)**\
         Returns documents that contain a specific prefix in a provided field.
  - [x] **[Range](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-range-query.html)**\
         Returns documents that contain terms within a provided range.
  - [ ] **[Regexp](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-regexp-query.html)**\
         Returns documents that contain terms matching a regular expression.
  - [x] **[Term](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-term-query.html)**\
         Returns documents that contain an exact term in a provided field.
  - [x] **[Terms](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-terms-query.html)**\
         Returns documents that contain one or more exact terms in a provided field.
  - [ ] **[Terms set](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-terms-set-query.html)**\
         Returns documents that contain a minimum number of exact terms in a provided field. You can define the minimum number of matching terms using a field or script.
  - [ ] **[Type](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-type-query.html) [Deprecated]**\
         Returns documents of the specified type.
  - [ ] **[Wildcard](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-wildcard-query.html)**\
         Returns documents that contain terms matching a wildcard pattern.
