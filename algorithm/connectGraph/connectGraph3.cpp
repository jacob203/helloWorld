class ConnectingGraph3 {
public:
    ConnectingGraph3(int n):mParentId(n+1) {
        for(int i = 1; i < mParentId.size(); i++){
            mParentId[i] = i;
            mGroups.insert(i);
        }
    }
        
    void  connect(int a, int b) {
        auto aLabel = Find(a);
        auto bLabel = Find(b);
        if(aLabel != bLabel){
            mParentId[aLabel] = bLabel;
            mGroups.erase(aLabel);
        }
    }

    int query() {
        return mGroups.size();
    }
    private:
        int Find(int i) {
            auto cur = i;
            while(i != mParentId[i])
            {
                i = mParentId[i];
            }
            mParentId[cur] = i;
            return i;
        }
    private:
        vector<int> mParentId;
        unordered_set<int> mGroups;
};