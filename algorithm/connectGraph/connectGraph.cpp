class ConnectingGraph {
public:
    ConnectingGraph(int n) {
        if (n <= 0) {
            return;//error
        }

        elements.resize(n+1);
        for(int i = 0; i < elements.size(); i++){
            elements[i] = i;
        }
    }

    void connect(int a, int b) {
        if( !isValidLabel(a) || !isValidLabel(b)) {
            return;
        }

        auto aParent = find(a);
        auto bParent = find(b);
        if (aParent != bParent) {
            elements[aParent] = bParent;
        }
    }

    bool query(int a, int b) {
        if (!isValidLabel(a) || !isValidLabel(b)){
            return false;
        }

        return find(a) == find(b);
    }
    private:
        vector<int> elements;
        bool isValidLabel(int label){
            if( label <= 0 || label > elements.size()) {
                return false;
            }
            return true;
        }

        int find(int index){
            while(elements[index] != index){
                index = elements[index];
            }
            return index;
        }
};

=======================================================================
compression find
class ConnectingGraph {
public:
    ConnectingGraph(int n) {
        if (n <= 0) {
            return;//error
        }

        elements.resize(n+1);
        for(int i = 0; i < elements.size(); i++){
            elements[i] = i;
        }
    }

    void connect(int a, int b) {
        if( !isValidLabel(a) || !isValidLabel(b)) {
            return;
        }

        auto aParent = findParent(a);
        auto bParent = findParent(b);
        if (aParent != bParent) {
            elements[aParent] = bParent;
        }
    }

    bool query(int a, int b) {
        if (!isValidLabel(a) || !isValidLabel(b)){
            return false;
        }

        return findParent(a) == findParent(b);
    }
    private:
        vector<int> elements;
        bool isValidLabel(int label){
            if( label <= 0 || label > elements.size()) {
                return false;
            }
            return true;
        }

        int findParent(int index) {
            auto curIndex = index;
            while(elements[index] != index){
                index = elements[index];
            }

            elements[curIndex] = index;
            return index;
        }
};
