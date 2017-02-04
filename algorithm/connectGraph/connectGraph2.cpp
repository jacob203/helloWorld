class ConnectingGraph2 {
public:
    ConnectingGraph2(int n):mParentId(n+1) {
        for(int i = 0; i < mParentId.size(); i++){
            mParentId[i] = make_pair(i, 1);
        }
    }

    void  connect(int a, int b) {
        if (a < 1 || a >= mParentId.size() || b < 1 || b >= mParentId.size()){
            return;//error
        }

        auto aId = UnionFind(a);
        auto bId = UnionFind(b);
        if(aId != bId) {
            mParentId[aId].first = bId;
            mParentId[bId].second += mParentId[aId].second;
        }
    }


    int query(int a) {
        if (a < 1 || a >= mParentId.size()){
            return -1;//error
        }

        auto aId = UnionFind(a);
        return mParentId[aId].second;
    }

    private:
    int UnionFind(int i) {
        while(i != mParentId[i].first){
            i = mParentId[i].first;
        }
        return i;
    }
    private:
        vector<pair<int,int>> mParentId;
};
========================================================================
class ConnectingGraph2 {
public:
    ConnectingGraph2(int n):mParentId(n+1) {
        for(int i = 0; i < mParentId.size(); i++){
            mParentId[i] = make_pair(i, 1);
        }
    }

    void  connect(int a, int b) {
        if (a < 1 || a >= mParentId.size() || b < 1 || b >= mParentId.size()){
            return;//error
        }

        auto aId = UnionFind(a);
        auto bId = UnionFind(b);
        if(aId != bId) {
            mParentId[aId].first = bId;
            mParentId[bId].second += mParentId[aId].second;
        }
    }


    int query(int a) {
        if (a < 1 || a >= mParentId.size()){
            return -1;//error
        }

        auto aId = UnionFind(a);
        return mParentId[aId].second;
    }

    private:
    int UnionFind(int i) {
        auto curId = i;
        while(i != mParentId[i].first){
            i = mParentId[i].first;
        }
        mParentId[curId].first = i;
        return i;
    }
    private:
        vector<pair<int,int>> mParentId;
};
