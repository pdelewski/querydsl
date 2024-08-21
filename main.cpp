#include <iostream>
#include <vector>
#include <tuple>

struct QueryDsl
{
  template<typename Table>
  Table from(Table t) {
    Table result;
    for(const auto& t : t.tuples) {
      result.tuples.push_back(t);
    }
    return result;
  }
  template<typename Table, typename Pred>
  Table where(Table t, Pred p)
  {
    Table result;
    for(const auto& t : t.tuples) {
      if(p(t)) {
        result.tuples.push_back(t);
      }
    }
    return result;    
  }

  template<typename OutTable, typename Table, typename Selector>
  OutTable select(Table t, Selector s)
  {
    OutTable result;
    for(const auto& t : t.tuples) {
      result.tuples.push_back(s(t));
    }
    return result;
  }

}; 

int main() {

  struct table
  {
    using tuple = struct{int field1; int field2;};
    std::vector<tuple> tuples;
    std::ostream& operator<<(std::ostream& os) {
      os << "eww";
      return os;
    }
  };

  struct table2
  {
    using tuple = struct{int field1;};
    std::vector<tuple> tuples;
  };


  QueryDsl q;
 
  auto t = table{
    {
     {1,1},
     {2,2}
    }
  };
  
  
  
  auto t1 = q.from(t);
  auto t2 = q.where(t1, [](const auto& r) {
    return r.field1 > 0;
  });
  auto t3 = q.select<table2>(t2, [](const auto& r) {
    return table2::tuple{r.field1};
  });
  for(const auto& t: t3.tuples) {
    std::cout << t.field1 << std::endl;
  }
  
  return 0;
}
