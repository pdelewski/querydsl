#include <iostream>
#include <vector>

struct QueryDsl {
  template <typename TableExpr> TableExpr from(TableExpr table_expr) {
    TableExpr result_tuple;
    for (const auto &tuple : table_expr.tuples) {
      result_tuple.tuples.push_back(tuple);
    }
    return result_tuple;
  }

  // selection
  template <typename TableExpr, typename Pred>
  TableExpr where(TableExpr table_expr, Pred p) {
    TableExpr result_tuple;
    for (const auto &tuple : table_expr.tuples) {
      if (p(tuple)) {
        result_tuple.tuples.push_back(tuple);
      }
    }
    return result_tuple;
  }

  // projection
  template <typename OutTableExpr, typename TableExpr, typename Selector>
  OutTableExpr select(TableExpr table_expr, Selector selector) {
    OutTableExpr result_tuple;
    for (const auto &tuple : table_expr.tuples) {
      result_tuple.tuples.push_back(selector(tuple));
    }
    return result_tuple;
  }
};

int main() {

  struct table {
    using tuple = struct {
      int field1;
      int field2;
    };
    std::vector<tuple> tuples;
  };

  struct result_table {
    using tuple = struct {
      int field1;
    };
    std::vector<tuple> tuples;
  };

  QueryDsl q;

  auto t0_expr = table{{{1, 1}, {2, 2}}};

  auto t1_expr = q.from(t0_expr);
  auto t2_expr =
      q.where(t1_expr, [](const auto &tuple) { return tuple.field1 > 0; });
  auto t3_expr = q.select<result_table>(t2_expr, [](const auto &tuple) {
    return result_table::tuple{tuple.field1};
  });

  for (const auto &tuple : t3_expr.tuples) {
    std::cout << tuple.field1 << std::endl;
  }

  return 0;
}
